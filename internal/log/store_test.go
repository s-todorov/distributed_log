package log

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogLoadAfterClose(t *testing.T) {
	dir := os.TempDir()

	store, err := NewIndexStore(dir)
	assert.NoError(t, err, "Error creating index store")

	tn := time.Now()

	for _, off := range store.indexes {
		at, err := store.ReadAt(off.Offset)
		assert.NoError(t, err, "Error reading from index store")
		fmt.Println(at)
	}

	store.Close()

	fmt.Println(len(store.indexes))
	fmt.Println(time.Since(tn))
}

func TestLogLoadAndWriteAfterClose(t *testing.T) {
	dir := os.TempDir()

	store, err := NewIndexStore(dir)
	assert.NoError(t, err, "Error creating index store")

	tn := time.Now()

	for _, off := range store.indexes {
		at, err := store.ReadAt(off.Offset)
		assert.NoError(t, err, "Error reading from index store")
		fmt.Println(at)
	}

	store.Close()

	store, err = NewIndexStore(dir)
	assert.NoError(t, err, "Error creating index store")

	_, err = store.Append(Index{File: fmt.Sprintf("test log:%d", 100), App: "test", State: "test", Msg: "test", ImportTime: time.Now(), Size: maxIndexSize})
	assert.NoError(t, err, "Error appending to index store")

	err = store.Sync()
	assert.NoError(t, err, "Error syncing index store")

	for _, off := range store.indexes {
		at, err := store.ReadAt(off.Offset)
		assert.NoError(t, err, "Error reading from index store")
		fmt.Println(at)
	}

	fmt.Println(time.Since(tn))
}

func TestLogReadWrite(t *testing.T) {
	dir := os.TempDir()
	l, err := NewIndexStore(dir)
	assert.NoError(t, err, "Error creating index store")

	tn := time.Now()
	for i := 0; i < 10; i++ {
		_, err = l.Append(Index{File: fmt.Sprintf("test log:%d", i), App: "test", State: "test", Msg: "test", ImportTime: time.Now(), Size: maxIndexSize})
		assert.NoError(t, err, "Error appending to index store")
	}

	l.Sync()

	for _, off := range l.indexes {
		at, err := l.ReadAt(off.Offset)
		assert.NoError(t, err, "Error reading from index store")
		fmt.Println(at)
	}

	err = l.Close()
	assert.NoError(t, err, "Error closing index store")

	fmt.Println(time.Since(tn))
}

func TestJson(t *testing.T) {
	dir := os.TempDir()

	storeFile, _ := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", 0, ".json")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)

	tn := time.Now()
	for i := 0; i < 100000; i++ {
		ind := Index{File: fmt.Sprintf("test log:%d", i), App: "test", State: "test", Msg: "test", ImportTime: time.Now(), Size: maxIndexSize}
		dt, err := json.Marshal(ind)
		assert.NoError(t, err, "Error marshaling index to JSON")

		_, err = storeFile.Write(dt)
		assert.NoError(t, err, "Error writing to store file")
	}
	fmt.Println(time.Since(tn))
}
