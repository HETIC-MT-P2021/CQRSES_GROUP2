package service

import (
	"cqrses/pkg/payload"
)

type Domain interface {
	NewPayload(status int, result interface{}) payload.DomainPayload
}

type DomainService struct{}

func (s DomainService) NewPayload(status int, result interface{}) payload.DomainPayload {
	return payload.NewPayload(status, result)
}
