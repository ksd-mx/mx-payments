package broker

type ProducerInterface[T any] interface {
	Publish(message T, key []byte, topic string) error
}
