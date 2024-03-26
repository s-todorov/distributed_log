package event

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"fmt"
	"testing"
	"time"
)

func TestNewHandler(t *testing.T) {
	m := mockIndexService{}
	h := NewHandler(&m)
	h.handlers[string(EventInsert)] = m.NewEvent

	e := Event{Type: string(EventInsert)}

	f := h.Run(&e)

	_, err := f(3)
	if err != nil {
		fmt.Println(err)
	}

	// publish/subscribe check
	newIndex := make(chan Event)
	newIndex2 := make(chan Event)

	go func() {
		res := <-newIndex
		fmt.Println("Subscriber 1", res.Type)

	}()
	go func() {
		res := <-newIndex2
		fmt.Println("Subscriber 2", res.Type)

	}()

	//h.EventBus.Subscribe(EventInsert, newIndex)
	//h.EventBus.Subscribe(EventInsert, newIndex2)
	// or
	go h.EventBus.Publish(&e)

	_, err = f(EventChan)
	if err != nil {
		fmt.Println(err)
	}

	h.EventBus.UnSubscribe(string(EventInsert))

	_, err = f(EventServ)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(3 * time.Second)
}

type mockIndexService struct {
	id int
}

var _ Indexer = new(mockIndexService)

func (m *mockIndexService) New() chan<- *Event {
	//TODO implement me
	panic("implement me")
}

func (m *mockIndexService) NewEvent(event *Event) (*v4.Record, error) {
	m.id++
	fmt.Println("hallo ", m.id)
	return nil, nil
}

// insert trough bus service
//process_through := ru.handler.Run(newEvent)
//err := process_through(3)
//if err != nil {
//	fmt.Println(err)
//}

// insert with publish/ All participant events will be triggered
//ru.handler.EventBus.Publish(newEvent)

// insert with service
//ru.handler.IndexService.New() <- newEvent

//ru.ser.Insert(&ni)
