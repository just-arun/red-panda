package main

import (
	"log"

	redpanda "github.com/just-arun/red-panda"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	client := redpanda.NewConnection(
		kgo.SeedBrokers("localhost:9092"),
		kgo.DefaultProduceTopic("foo"),
		kgo.ConsumeTopics("foo"),
	)
	if err := redpanda.RedPanda(client).Publish("foo", []byte("hello world")); err != nil {
		log.Fatal(err)
	}
}
