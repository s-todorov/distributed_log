package log

import (
	"bufio"
	"bytes"
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/go-multierror"
	"io"
	"log/slog"
	"os"
	"path"
	"sync"
)

// TODO check and remove dead code
//type Loger interface {
//	ReadAt(offset int64) (index *indx.Index, err error)
//	Append(index indx.Index) (offset int64, err error)
//	Close() error
//}

// Log represents the main structure for storing and managing indexes.
type Log struct {
	mu                sync.RWMutex
	file              *os.File
	bufWriter         *bufio.Writer
	Dir               string
	currSegmentOffset int64
	segment           *segment
	indexes           []*storeInMemory
	Config            Config
	last, first       int64
}

func (s *Log) Reader() io.Reader {
	s.mu.RLock()
	defer s.mu.RUnlock()
	//readers := make([]io.Reader, len(s.segments))
	//for i, segment := range s.segments {
	//	readers[i] = &originReader{segment.File, 0}
	//}
	//return io.MultiReader(readers...)
	return s.segment.file
}

// storeInMemory represents information about the stored indexes in memory.
type LogInMemory struct {
	Offset int64
	File   string
}

// NewIndexStore creates a new instance of IndexStore.
func NewLog(dir string) (*Log, error) {
	baseOffset := 1

	storeDir := path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".store"))
	file, err := os.OpenFile(
		storeDir,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)

	s, err := newSegment(dir)
	if err != nil {
		return nil, err
	}

	store := Log{
		file:      file,
		segment:   s,
		indexes:   make([]*storeInMemory, 0),
		bufWriter: bufio.NewWriter(file),
		Dir:       dir,
	}
	return &store, store.setup()
}
func (s *Log) AddRaft(c Config) {
	s.Config = c
}

// setup initializes the IndexStore by reading existing indexes from the file.
func (s *Log) setup() error {
	info, err := s.file.Stat()
	if err != nil {
		return err
	}
	last := info.Size()
	s.currSegmentOffset = s.segment.size

	for {
		at, off, err := readAt(s.file, last)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		dec, err := s.dec(at)
		if err != nil {
			return err
		}
		s.indexes = append(s.indexes, dec)
		last = off

		if off == 0 {
			break
		}
	}

	return nil
}

// Append adds a new index to the store.
func (s *Log) Append(index *v4.Record) (offset int64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer func(s *Log) {
		err := s.Sync()
		if err != nil {

		}
	}(s)

	bt, err := proto.Marshal(index)
	wbt, err := s.segment.write(bt)
	if err != nil {
		return 0, err
	}
	s.currSegmentOffset += int64(wbt)

	storeMem := &storeInMemory{File: s.segment.Name(), Offset: s.currSegmentOffset}
	s.indexes = append(s.indexes, storeMem)

	data, err := s.enc(storeMem)
	if err != nil {
		return 0, err
	}

	_, err = s.bufWriter.Write(data)
	if err != nil {
		return 0, err
	}
	off := make([]byte, offWidth)
	binary.LittleEndian.PutUint32(off, uint32(len(data)))

	_, err = s.bufWriter.Write(off)
	if err != nil {
		return 0, err
	}

	return s.currSegmentOffset, nil
}

// ReadAt retrieves an index from the specified offset.
func (s *Log) ReadAt(offset int64) (index *v4.Record, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i, _, err := s.segment.popAt(offset)
	if err != nil {
		return nil, err
	}

	ind := &v4.Record{}
	err = proto.Unmarshal(i, ind)
	if err != nil {
		return nil, err
	}

	return ind, nil
}

func (s *Log) Read(offset int64) ([]byte, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.segment.popAt(offset)
}

func (l *Log) LowestOffset() (uint64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if len(l.indexes) <= 0 {
		return 0, nil
	}
	//return uint64(l.indexes[0].Offset), nil
	return uint64(len(l.indexes)), nil
}

func (l *Log) HighestOffset() (uint64, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if len(l.indexes) <= 0 {
		return 0, nil
	}

	return uint64(len(l.indexes) - 1), nil
}

// Close closes the IndexStore and flushes/syncs the segments and store.
func (s *Log) Close() error {
	var closeErr error

	slog.Info("flush segment")
	err := s.segment.bufWriter.Flush()
	if err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	slog.Info("sync segment")
	err = s.segment.file.Sync()
	if err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	slog.Info("flush store")
	err = s.bufWriter.Flush()
	if err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	slog.Info("sync store")
	err = s.file.Sync()
	if err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	slog.Info("close segment")
	if err := s.segment.file.Close(); err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	slog.Info("close store")
	if err := s.file.Close(); err != nil {
		closeErr = multierror.Append(closeErr, err)
	}
	return closeErr
}

// Sync flushes and syncs the segments and store.
func (s *Log) Sync() error {
	var syncErr error

	err := s.segment.bufWriter.Flush()
	if err != nil {
		syncErr = multierror.Append(syncErr, err)
	}
	err = s.bufWriter.Flush()
	if err != nil {
		syncErr = multierror.Append(syncErr, err)
	}
	err = s.segment.file.Sync()
	if err != nil {
		syncErr = multierror.Append(syncErr, err)
	}
	err = s.file.Sync()
	if err != nil {
		syncErr = multierror.Append(syncErr, err)
	}
	return syncErr
}

// Remove closes and removes the store directory.
func (s *Log) Remove() error {
	if err := s.segment.file.Close(); err != nil {
		return err
	}
	return os.RemoveAll(s.Dir)
}

// Reset removes the store directory and sets up a new store.
func (s *Log) Reset() error {
	if err := s.Remove(); err != nil {
		return err
	}
	//return s.setup()
	return nil
}

// enc serializes storeInMemory to bytes using gob.
func (s *Log) enc(index *storeInMemory) (dt []byte, err error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err = enc.Encode(index)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}

// dec deserializes bytes to storeInMemory using gob.
func (s *Log) dec(idt []byte) (index *storeInMemory, err error) {
	var network bytes.Buffer

	_, err = network.Write(idt)
	if err != nil {
		return nil, fmt.Errorf("(s *IndexStore) dec => network.Write(idt): %w", err)
	}
	dec := gob.NewDecoder(&network)

	var i storeInMemory
	err = dec.Decode(&i)
	if err != nil {
		return nil, fmt.Errorf("(s *IndexStore) dec => dec.Decode(&i): %w", err)
	}

	return &i, nil
}
