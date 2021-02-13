package cqrs

import (
	"fmt"
	"reflect"
)

// CommandMessage define the structure of the command message entity
type CommandMessage interface {
	CommandType() string
	Payload() interface{}
}

// CommandBus define the structure of the command bus entity
type CommandBus struct {
	handlers map[string]CommandHandler
}

// CommandDescriptor define the structure of the command descriptor entity
type CommandDescriptor struct {
	command interface{}
}

// NewCommandBus return a new command bus
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler),
	}
}

// Dispatch dispatch the commands
func (b *CommandBus) Dispatch(command CommandMessage) (interface{}, error) {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}

	return nil, fmt.Errorf("fatal error command: The command bus does not have a handler for commands of type: %s", command.CommandType())
}

// RegisterHandler register the command handler
func (b *CommandBus) RegisterHandler(handler CommandHandler, command interface{}) error {
	typeName := reflect.TypeOf(command).Elem().Name()
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("fatal error command: Duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

// NewCommandMessage create a new command message
func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

// CommandType return the command type
func (c *CommandDescriptor) CommandType() string {
	return reflect.TypeOf(c.command).Elem().Name()
}

// Payload returns the actual command payload
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}
