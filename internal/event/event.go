package event

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"github.com/pkg/errors"
	"time"
)

type Event struct {
	Type      string
	Timestamp time.Time
	Data      interface{}
}

type MessageType int32

const (
	EventChan MessageType = 0
	EventServ MessageType = 1
)

type eventType string

const (
	EventInsert eventType = "insert"
	EventGet    eventType = "get"
)

var (
	ErrEventNotSupported         = errors.New("this event type is not supported")
	ErrEventNoChanSubscribers    = errors.New("no Chan subscribers are added")
	ErrEventNoServiceSubscribers = errors.New("no Service subscribers are added")
)

type Indexer interface {
	NewEvent(event *Event) (*v4.Record, error)
	New() chan<- *Event
}

type Store interface {
	Read(offset uint64) (*v4.Record, error)
	Append(record *v4.Record) (uint64, error)
}
