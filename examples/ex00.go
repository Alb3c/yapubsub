package main

import(
	"fmt"
	"yapubsub"
)

func main() {
	b := yapubsub.NewBroker()
	s1 := b.Subscribe("topic001")
	s2 := b.Subscribe("topic001")

	go (func(s *yapubsub.Subscriber) {
		for {
			msg := s.WaitMessage()
			fmt.Printf("Func1: Found new message: %s\n", msg)
		}
	})(s1)

	go (func(s *yapubsub.Subscriber) {
		for {
			msg := s.WaitMessage()
			fmt.Printf("Func2: Found new message: %s\n", msg)
		}
	})(s2)
	
	b.Publish("topic001", "This is a test")
	b.Unsubscribe("topic001", s1)
	b.Publish("topic001", "Another test")

	p := b.Publisher("topic001")
	p.Publish("Test with Publisher")
}