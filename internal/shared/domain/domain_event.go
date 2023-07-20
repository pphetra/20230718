package shared_domain

type DomainEvent interface {
	GetName() string
	GetPayload() []byte
}
