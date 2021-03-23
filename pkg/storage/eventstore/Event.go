package eventstore

import (
	"cqrses/pkg/storage"
	"time"
)

// Event defines the structure of the events to be stored
type Event struct {
	EventName      string                `json:"event_name"`
	EventType      EventTypology         `json:"event_type"`
	ObjectID       string                `json:"object_id"`
	ObjectVersion  storage.ObjectVersion `json:"object_version"`
	Payload        interface{}           `json:"payload"`
	EventCreatedAt time.Time             `json:"event_created_at"`
}

// EventTypology of an event
type EventTypology string

// typology types
const (
	Create EventTypology = "create"
	Update EventTypology = "update"
	Delete EventTypology = "delete"
)
