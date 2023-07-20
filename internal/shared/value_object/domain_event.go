package value_object

type DomainEvent interface {
	GetName() string
	GetPayload() []byte
}
