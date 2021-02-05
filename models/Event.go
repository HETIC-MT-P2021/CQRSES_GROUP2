package models

import (
	"time"
)

// Event define the structure of the event element
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
