package event

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"distributed_log/internal/log"
	"fmt"
	"log/slog"
	"strconv"
)

// EventBus represents the event bus that handles event subscription and dispatching
type EventBus struct {
	subscribers        map[string][]chan<- *Event
	serviceSubscribers map[string][]EventHandler
	store              Store
	nodes              []*log.DistributedLog
}

// NewEventBus creates a new instance of the event bus
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan<- *Event),
	}
}

// Subscribe adds a new subscriber for a given event type
// TODO check if exist before append
func (eb *EventBus) Subscribe(eventType string, indexer Indexer) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], indexer.New())
}

func (eb *EventBus) SubscribeService(eventType string, subscriber EventHandler) {
	eb.serviceSubscribers[eventType] = append(eb.serviceSubscribers[eventType], subscriber)
}

// UnSubscribe remove the subscriber for a given event type
// TODO if not exist return warning
func (eb *EventBus) UnSubscribe(eventType string) {
	delete(eb.subscribers, eventType)
}

// Publish sends an event to all subscribers of a given event type
func (eb *EventBus) Publish(event *Event) (*v4.Record, error) {

	offset := 0
	if event.Type == string(EventInsert) {
		_, leaderID := eb.nodes[0].GetLeader()

		for _, node := range eb.nodes {

			if node.LogID == string(leaderID) {
				if pr, ok := event.Data.(*v4.ProduceRequest); ok {
					slog.Info("ewIndexProto, ok := event.Data.(*v4.Record); ok OK")
					slog.Info(strconv.FormatUint(pr.Record.Offset, 10))
					off, err := node.Append(pr.Record)
					if err != nil {
						return nil, err
					}
					offset = int(off)
				} else {
					slog.Info("ewIndexProto, ok := event.Data.(*v4.Record); ok NOK")
					slog.Info(event.Type)
				}

				eb.store = node
				break
			}
		}

		//o, err := eb.appendToLog(event)
		//if err != nil {
		//	return nil, err
		//}
		//off = int(o)
	}
	if event.Type == string(EventGet) {
		indx, err := eb.read(event)
		if err != nil {
			return nil, err
		}
		return indx, nil
	}

	subscribers := eb.subscribers[event.Type]
	for _, subscriber := range subscribers {
		subscriber <- event
	}
	return &v4.Record{Offset: uint64(offset)}, nil
}

// Stop close all subscribers chan
func (eb *EventBus) Stop() {

	for t, subscriber := range eb.subscribers {
		eventType := t
		for _, v := range subscriber {
			slog.Info(fmt.Sprintf("close event type:%s", eventType))
			close(v)
		}
	}
}

// Publish sends an event to all subscribers of a given event type
func (eb *EventBus) publishService(event *Event) {
	subscribers := eb.serviceSubscribers[event.Type]
	for _, subscriber := range subscribers {
		_, err := subscriber(event)
		if err != nil {
			return
		}
	}
}

// HasSubscriber check if subscriber have been added to the event bus
func (eb *EventBus) HasSubscriber(eventType string) bool {
	if _, ok := eb.subscribers[eventType]; ok {
		return true
	}
	return false
}

// HasServiceSubscriber check if subscriber have been added to the event bus
func (eb *EventBus) HasServiceSubscriber(eventType string) bool {
	if _, ok := eb.serviceSubscribers[eventType]; ok {
		return true
	}
	return false
}

func (eb *EventBus) appendToLog(event *Event) (int64, error) {
	if eb.store != nil {
		pr, ok := event.Data.(*v4.ProduceRequest)
		if ok && pr != nil {
			off, err := eb.store.Append(pr.Record)
			if err != nil {
				return -1, fmt.Errorf("(eb *EventBus) Publish =>store.Append: %w", err)
			}
			return int64(off), nil
		}
	}
	return -1, fmt.Errorf("store is == nil")
}

func (eb *EventBus) read(event *Event) (*v4.Record, error) {
	if eb.store != nil {
		pr, ok := event.Data.(*v4.ConsumeRequest)
		if ok && pr != nil {
			off, err := eb.store.Read(pr.Offset)
			if err != nil {
				return nil, fmt.Errorf("(eb *EventBus) Publish =>store.ReadAt: %w", err)
			}
			return off, nil
		}
	}
	return nil, fmt.Errorf("store is == nil")
}
