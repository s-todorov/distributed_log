package log

import (
	"bufio"
	"bytes"
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// segment represents a segment for storing indexes.
type segment struct {
	file      *os.File
	size      int64
	bufWriter *bufio.Writer
	bufReader *bufio.Reader
}

// newSegment creates a new instance of the segment.
func newSegment(dir string) (*segment, error) {
	baseOffset := 1

	file, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".seg")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)

	idx := &segment{
		file:      file,
		bufWriter: bufio.NewWriter(file),
		bufReader: bufio.NewReader(file),
	}
	fi, err := os.Stat(file.Name())
	if err != nil {
		return nil, err
	}

	idx.size = fi.Size()
	return idx, nil
}

// write writes the given data to the segment file.
func (s *segment) write(p []byte) (int, error) {
	if len(p) > maxIndexSize {
		return 0, errors.New("the index is too large")
	}

	// todo remove last write if err
	if err := binary.Write(s.bufWriter, enc, p); err != nil {
		return 0, err
	}

	off := make([]byte, offWidth)
	binary.LittleEndian.PutUint32(off, uint32(len(p)))

	w, err := s.bufWriter.Write(off)
	if err != nil {
		return w, err
	}
	writtenOffset := w + len(p)

	return writtenOffset, nil
}

// popLast removes and retrieves the last index from the segment.
func (s *segment) popLast() (index []byte, err error) {
	fileInfo, err := s.file.Stat()
	if err != nil {
		return nil, err
	}
	lastPosition := fileInfo.Size() - offWidth
	indexSize := make([]byte, offWidth)
	_, err = s.file.ReadAt(indexSize, lastPosition)
	if err != nil {
		return nil, err
	}

	byteToRead := binary.LittleEndian.Uint32(indexSize)
	dataSize := fileInfo.Size() - int64(byteToRead) - offWidth

	if fileInfo.Size() == int64(byteToRead)+offWidth {
		dataSize = 0
	}

	size := make([]byte, byteToRead)
	_, err = s.file.ReadAt(size, dataSize)
	if err != nil {
		return nil, err
	}
	return size, nil
}

// popAt removes and retrieves the index at the specified offset.
func (s *segment) popAt(offset int64) (index []byte, readData int, err error) {
	indexSize := make([]byte, offWidth)

	_, err = s.file.ReadAt(indexSize, offset-int64(offWidth))
	if err != nil {
		return nil, 0, fmt.Errorf("error popAt => s.file.ReadAt(indexSize, offset-offWidth): %w", err)
	}

	byteToRead := binary.LittleEndian.Uint32(indexSize)
	dataSize := offset - int64(byteToRead) - int64(offWidth)

	size := make([]byte, byteToRead)
	rd, err := s.file.ReadAt(size, dataSize)
	if err != nil {
		return nil, 0, fmt.Errorf("error popAt => i.File.ReadAt(size, dataSize): %w", err)
	}
	return size, rd, nil
}

// Name returns the name of the segment file.
func (s *segment) Name() string {
	return s.file.Name()
}

// encProto serializes an index to bytes using protobuf.
func (s *segment) encProto(index Index) ([]byte, error) {
	idnex := &v4.Record{Offset: index.Offset, Value: []byte(index.File)}
	marshal, err := proto.Marshal(idnex)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

// decProto deserializes bytes to an index using protobuf.
func (s *segment) decProto(idt []byte) (*Index, error) {
	ind := &v4.Record{}
	err := proto.Unmarshal(idt, ind)
	if err != nil {
		fmt.Print(err)
	}

	return &Index{ImportTime: time.Now(), File: string(ind.Value)}, nil
}

// enc serializes an index to bytes using gob.
func (s *segment) enc(index Index) (dt []byte, err error) {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err = enc.Encode(index)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}

// dec deserializes bytes to an index using gob.
func (s *segment) dec(idt []byte) (index *Index, err error) {
	var network bytes.Buffer

	_, err = network.Write(idt)
	if err != nil {
		return nil, fmt.Errorf("(s *IndexStore) dec => network.Write(idt): %w", err)
	}
	dec := gob.NewDecoder(&network)

	var i Index
	err = dec.Decode(&i)
	if err != nil {
		return nil, fmt.Errorf("(s *IndexStore) dec => dec.Decode(&i): %w", err)
	}

	return &i, nil
}
