package eventstore

import (
	"cqrses/pkg/storage"
	"time"
)

// Event defines the structure of the events to be stored
type Event struct {
	EventName      string
	EventType      EventTypology
	ObjectID       string
	ObjectVersion  storage.ObjectVersion
	Payload        interface{}
	EventCreatedAt time.Time
}

// EventTypology of an event
type EventTypology string

// typology types
const (
	Create EventTypology = "create"
	Update EventTypology = "update"
	Delete EventTypology = "delete"
)
