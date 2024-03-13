package log

import "time"

type Index struct {
	File       string
	App        string
	State      string
	Msg        string
	ImportTime time.Time
	Size       int64
	Offset     uint64
}

type SegmentIndex struct {
	Offset uint64
}

func NewIndex(file string, app string, state string, msg string, size int64) Index {

	return Index{File: file, App: app, State: state, Msg: msg, ImportTime: time.Now(), Size: size}
}
