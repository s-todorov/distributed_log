package service

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"distributed_log/internal/event"
	"fmt"
	"log/slog"
)

type IndexService struct {
	rep event.Store
	new chan<- *event.Event
}

func (ap *IndexService) New() chan<- *event.Event {
	return ap.new
}

func (ap *IndexService) NewEvent(e *event.Event) (*v4.Record, error) {
	if e.Type == string(event.EventInsert) {
		off, err := ap.Insert(e)
		if err != nil {
			return nil, err
		}
		return &v4.Record{Offset: off}, nil
	}
	if e.Type == string(event.EventGet) {
		return ap.Get(e)
	}
	// todo return err
	return nil, nil
}

// ADD READ AND APPEN INDEXER INTERFACE
func (ap *IndexService) Insert(event *event.Event) (uint64, error) {

	switch v := event.Data.(type) {
	case *v4.Record:
		return ap.rep.Append(v)
	default:
		slog.Error("newEvent invalid event data")
		return 0, fmt.Errorf("invalid event data")
	}

}

func (ap *IndexService) Get(event *event.Event) (*v4.Record, error) {
	switch v := event.Data.(type) {
	case *v4.Record:
		return ap.rep.Read(v.Offset)
	default:
		slog.Error("newEvent invalid event data")
		return nil, fmt.Errorf("invalid event data")
	}
}

func NewIndexService(repository event.Store) *IndexService {

	indexService := IndexService{rep: repository}

	return &indexService
}
