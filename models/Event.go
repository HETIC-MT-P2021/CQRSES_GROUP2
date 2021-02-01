package models

import (
	"time"
)

// Event define the structure of the event element
type Event struct {
	ID        string
	Typology  Typology
	Payload   interface{}
	CreatedAt time.Time
	Index     uint
}

// Typology of an event
type Typology string

// typology types
const (
	Create Typology = "create"
	Update Typology = "update"
	Delete Typology = "delete"
)
