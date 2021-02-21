package service

type Domain interface {
	NewPayload(status int, result interface{}) DomainPayload
}

type DomainService struct{}

func (s DomainService) NewPayload(status int, result interface{}) DomainPayload {
	return NewPayload(status, result)
}
