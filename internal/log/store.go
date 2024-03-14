package log

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/raft"
	"io"
	"log/slog"
	"os"
	"path"
	"sync"
)

var (
	enc = binary.LittleEndian
)

const (
	maxIndexSize = 1024
	offWidth     = 4
)

// IndexStore represents the main structure for storing and managing indexes.
type IndexStore struct {
	mu                sync.RWMutex
	file              *os.File
	bufWriter         *bufio.Writer
	Dir               string
	currSegmentOffset int64
	segment           *segment
	indexes           []*storeInMemory
	Config            Config
}

func (s *IndexStore) Reader() io.Reader {
	s.mu.RLock()
	defer s.mu.RUnlock()
	//readers := make([]io.Reader, len(s.segments))
	//for i, segment := range s.segments {
	//	readers[i] = &originReader{segment.File, 0}
	//}
	//return io.MultiReader(readers...)
	return s.segment.file
}

type originReader struct {
	*storeInMemory
	off int64
}

//func (o *originReader) Read(p []byte) (int, error) {
//	n, err := o.ReadAt(p, o.off)
//	o.off += int64(n)
//	return n, err
//}

func (s *IndexStore) FirstIndex() (uint64, error) {
	return 0, nil
}

func (s *IndexStore) LastIndex() (uint64, error) {
	return 0, nil
}

func (s *IndexStore) GetLog(index uint64, log *raft.Log) error {
	return nil
}

func (s *IndexStore) StoreLog(log *raft.Log) error {
	return nil
}

func (s *IndexStore) StoreLogs(logs []*raft.Log) error {
	return nil
}

func (s *IndexStore) DeleteRange(min, max uint64) error {
	return nil
}

// storeInMemory represents information about the stored indexes in memory.
type storeInMemory struct {
	Offset int64
	File   string
}

// NewIndexStore creates a new instance of IndexStore.
func NewIndexStore(dir string) (*IndexStore, error) {
	baseOffset := 1

	storePath := path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".store"))
	//err := os.Mkdir(storePath, 0644)
	//if err != nil {
	//	return nil, err
	//}

	file, err := os.OpenFile(
		storePath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}

	log, err := newSegment(dir)
	if err != nil {
		return nil, err
	}

	store := IndexStore{
		file:      file,
		segment:   log,
		indexes:   make([]*storeInMemory, 0),
		bufWriter: bufio.NewWriter(file),
	}
	return &store, store.setup()
}
func (s *IndexStore) AddRaft(c Config) {
	s.Config = c
}

// setup initializes the IndexStore by reading existing indexes from the file.
func (s *IndexStore) setup() error {
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

// readAt reads data from a specified offset in the file.
func readAt(f *os.File, pos int64) ([]byte, int64, error) {
	offset := pos - offWidth
	if offset <= -1 {
		return nil, 0, io.EOF
	}

	segmentSize := make([]byte, offWidth)
	_, err := f.ReadAt(segmentSize, offset)
	if err != nil {
		return nil, 0, err
	}

	byteToRead := binary.LittleEndian.Uint32(segmentSize)
	dataSize := pos - int64(byteToRead) - offWidth

	size := make([]byte, byteToRead)
	_, err = f.ReadAt(size, dataSize)
	if err != nil {
		return nil, 0, err
	}
	return size, dataSize, nil
}

// Append adds a new index to the store.
func (s *IndexStore) Append(index Index) (offset int64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	defer func(s *IndexStore) {
		// todo flush to disk after time or size
		err := s.Sync()
		if err != nil {

		}
	}(s)

	bt, err := s.segment.encProto(index)
	if err != nil {
		return 0, err
	}

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
func (s *IndexStore) ReadAt(offset int64) (index *Index, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	i, _, err := s.segment.popAt(offset)
	if err != nil {
		return nil, err
	}
	return s.segment.decProto(i)
}

// PopLast removes and retrieves the last index from the store.
func (s *IndexStore) PopLast() (index *Index, err error) {
	i, err := s.segment.popLast()
	if err != nil {
		return nil, err
	}
	return s.segment.dec(i)
}

// Close closes the IndexStore and flushes/syncs the segments and store.
func (s *IndexStore) Close() error {
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
func (s *IndexStore) Sync() error {
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
func (s *IndexStore) Remove() error {
	if err := s.segment.file.Close(); err != nil {
		return err
	}
	return os.RemoveAll(s.Dir)
}

// Reset removes the store directory and sets up a new store.
func (s *IndexStore) Reset() error {
	if err := s.Remove(); err != nil {
		return err
	}
	//return s.setup()
	return nil
}

// enc serializes storeInMemory to bytes using gob.
func (s *IndexStore) enc(index *storeInMemory) (dt []byte, err error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err = enc.Encode(index)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}

// dec deserializes bytes to storeInMemory using gob.
func (s *IndexStore) dec(idt []byte) (index *storeInMemory, err error) {
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
