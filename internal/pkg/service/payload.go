package service

type DomainPayload interface {
	GetStatus() int         // Returns the domain status.
	GetResult() interface{} // Returns the result produced by the domain.
}

type Payload struct {
	Status int
	Result interface{}
}

// NewPayload create a new Payload instance.
func NewPayload(status int, result interface{}) DomainPayload {
	p := new(Payload)
	p.Status = status
	p.Result = result

	return p
}

// GetStatus returns the domain status.
func (p Payload) GetStatus() int {
	return p.Status
}

// GetResult returns the result produced by the domain.
func (p Payload) GetResult() interface{} {
	return p.Result
}
