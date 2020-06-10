package yapubsub

import(
	"sync"
	"errors"
)

type Message struct {
	Data interface{}
}

type Broker struct {
	topics map[string][]*Subscriber
	lock  sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string][]*Subscriber),
	}
}

func (b *Broker) Subscribe(topic string) *Subscriber {
	b.lock.Lock()
	defer b.lock.Unlock()

	if _, ok := b.topics[topic]; !ok {
		b.topics[topic] = []*Subscriber{}
	}
	s := NewSubscriber(topic)
	b.topics[topic] = append(b.topics[topic], s)
	return s
}

func (b *Broker) Unsubscribe(topic string, sub *Subscriber) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if _, ok := b.topics[topic]; !ok {
		return errors.New("Unable unsubscribe: Topic not valid")
	}
	for i, s := range b.topics[topic] {
		if s == sub {
			/* Remove element */
			b.topics[topic] = append(b.topics[topic][:i], b.topics[topic][i+1:]...)
			break
		}
	}
	return nil
}

func (b *Broker) Publisher(topic string) *Publisher {
	return NewPublisher(topic, b)
}

func (b *Broker) Publish(topic string, m interface{}) error {
	if _, ok := b.topics[topic]; !ok {
		return errors.New("Unable to publish message: No subscriber have been found")
	}
	/* Send the message over all the subscribers */
	msg := &Message{
		Data: m,
	}
	for _, s := range b.topics[topic] {
		s.msgs <- msg
	}
	return nil
}