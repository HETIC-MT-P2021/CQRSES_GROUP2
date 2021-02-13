package eventstore

import (
	"time"
)

// Event defines the structure of the events to be stored
type Event struct {
	Name      string
	Typology  Typology
	ObjectID  string
	Payload   interface{}
	CreatedAt time.Time
}

// Typology of an event
type Typology string

// typology types
const (
	Create Typology = "create"
	Update Typology = "update"
	Delete Typology = "delete"
)
