package main

import (
	"context"
	"fmt"
	"log"

	redpanda "github.com/just-arun/red-panda"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	client := redpanda.NewConnection(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("foo"),
	)
	redpanda.RedPanda(client).Subscribe(context.Background(), "foo", func(err error, msg *kgo.Record) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg.Value))
	})
}
