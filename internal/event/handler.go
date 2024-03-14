package event

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"distributed_log/internal/log"
)

type Handler struct {
	hostname       string
	releaseVersion string

	handlers     map[string]EventHandler
	EventBus     *EventBus
	IndexService Indexer
}

type EventHandler func(event *Event) (*v4.Record, error)

func NewHandler(indexer Indexer, ev ...eventType) *Handler {
	cli := Handler{handlers: make(map[string]EventHandler), EventBus: NewEventBus(), IndexService: indexer}

	for _, evn := range ev {
		cli.EventBus.Subscribe(string(evn), indexer)
	}

	cli.setupEventServiceHandlers()
	return &cli
}
func (c *Handler) WithStoreLog(store []*log.DistributedLog) *Handler {
	c.EventBus.nodes = store
	return c
}

func (c *Handler) Stop() error {
	//return c.EventBus.store.Close()
	return nil
}

// TODO MOVE setup to another layer
// setupEventServiceHandlers initializes the event handlers in the Handler struct.
// It assigns the appropriate Handler functions to specific event types.
func (c *Handler) setupEventServiceHandlers() {
	c.handlers[string(EventInsert)] = c.IndexService.NewEvent
	c.handlers[string(EventGet)] = c.IndexService.NewEvent
}

func (c *Handler) routeEventChan(event *Event) (*v4.Record, error) {

	// if we have added more than one subscriber, trigger all of them, otherwise only execute the service function
	if c.EventBus.HasSubscriber(event.Type) {
		return c.EventBus.Publish(event)

	}
	return nil, ErrEventNoChanSubscribers

}

// routeEvent routes the given event to the corresponding Handler based on the event type.
// It retrieves the Handler function associated with the event type from the handlers map
// and executes it with the provided event, cli, and cobra arguments.
// If the event type is not supported, it returns ErrEventNotSupported.
func (c *Handler) routeEventService(event *Event) (*v4.Record, error) {

	// if we have added more than one subscriber, trigger all of them, otherwise only execute the service function
	if c.EventBus.HasServiceSubscriber(event.Type) {
		c.EventBus.publishService(event)
	}
	return nil, ErrEventNoServiceSubscribers
}

// routeEvent routes the given event to the corresponding Handler based on the event type.
// It retrieves the Handler function associated with the event type from the handlers map
// and executes it with the provided event, cli, and cobra arguments.
// If the event type is not supported, it returns ErrEventNotSupported.
func (c *Handler) routeEvent(event *Event) (*v4.Record, error) {
	if h, ok := c.handlers[event.Type]; ok {
		l, err := h(event)
		if err != nil {
			return nil, err
		}
		return l, nil
	} else {
		return nil, ErrEventNotSupported
	}
}

// Run returns a function that can be used as the RunE field in a Cobra command.
// The returned function routes the provided event to the appropriate handler using routeEvent.
func (c *Handler) Run(event *Event) func(args MessageType) (*v4.Record, error) {
	return func(args MessageType) (*v4.Record, error) {

		switch args {
		case EventChan:
			return c.routeEventChan(event)
		case EventServ:
			// todo return offset
			return c.routeEventService(event)
		default:
			// todo return offset
			return c.routeEvent(event)
		}
	}
}
