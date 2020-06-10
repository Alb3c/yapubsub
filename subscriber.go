package yapubsub

import (
)

type Subscriber struct {
	topic 	string
	msgs  	chan *Message
}

func NewSubscriber(t string) *Subscriber {
	return &Subscriber{
		topic: t,
		msgs: make(chan *Message),
	}
}

func (s *Subscriber) GetTopic() string {
	return s.topic
}

func (s *Subscriber) WaitMessage() interface{} {
	m := <-s.msgs
	return *m
}