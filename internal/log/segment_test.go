package log

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSegment(t *testing.T) {
	dir := os.TempDir()

	s, err := newSegment(dir)
	assert.NoError(t, err, "Error creating")

	data := []byte("Hallo")
	_, err = s.write(data)
	assert.NoError(t, err, "Error writing")

	err = s.bufWriter.Flush()
	assert.NoError(t, err, "Error flush")

	err = s.file.Sync()
	assert.NoError(t, err, "Error sync")

	last, err := s.popLast()
	assert.NoError(t, err, "Error popLast")

	assert.Equal(t, data, last)
}
