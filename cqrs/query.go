package cqrs

import (
	"fmt"
	"reflect"
)

// QueryMessage define the structure of query message entity
type QueryMessage interface {
	Payload() interface{}
	QueryType() string
}

// QueryBus define the structure of query bus entity
type QueryBus struct {
	handlers map[string]QueryHandler
}

// QueryDescriptor define the structure of query descriptor entity
type QueryDescriptor struct {
	query interface{}
}

// NewQueryBus create a new query bus
func NewQueryBus() *QueryBus {
	cBus := &QueryBus{
		handlers: make(map[string]QueryHandler),
	}

	return cBus
}

// Dispatch dispatch the queries
func (b *QueryBus) Dispatch(query QueryMessage) (interface{}, error) {
	if handler, ok := b.handlers[query.QueryType()]; ok {
		return handler.Handle(query)
	}

	return nil, fmt.Errorf("fatal error query: The query bus does not have a handler for query of type: %s", query.QueryType())
}

// RegisterHandler register the query handler
func (b *QueryBus) RegisterHandler(handler QueryHandler, query interface{}) error {
	typeName := reflect.TypeOf(query).Elem().Name()
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("fatal error query: Duplicate query handler registration with query bus for query of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

// NewQueryMessage create a new query message
func NewQueryMessage(query interface{}) *QueryDescriptor {
	return &QueryDescriptor{
		query: query,
	}
}

// QueryType return the type of query
func (c *QueryDescriptor) QueryType() string {
	return reflect.TypeOf(c.query).Elem().Name()
}

// Payload returns the actual query payload
func (c *QueryDescriptor) Payload() interface{} {
	return c.query
}
