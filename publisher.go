package yapubsub

import (
)

type Publisher struct {
	topic 	string
	broker	*Broker
}

func NewPublisher(t string, b *Broker) *Publisher {
	return &Publisher{
		topic: t,
		broker: b,
	}
}

func (p *Publisher) GetTopic() string {
	return p.topic
}

func (p *Publisher) Publish(m interface{}) error {
	return p.broker.Publish(p.topic, m)
}