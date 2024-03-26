package log

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSegment(t *testing.T) {
	dir := os.TempDir()

	s, err := newSegment(dir)
	assert.NoError(t, err, "Error creating index store")

	off, err := s.write([]byte("Hallo"))
	assert.NoError(t, err, "Error creating index store")
	fmt.Println(off)
}
