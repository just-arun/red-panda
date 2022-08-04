package redpanda

import "github.com/twmb/franz-go/pkg/kgo"

type redPanda struct {
	*kgo.Client
}

func RedPanda(client *kgo.Client) *redPanda {
	return &redPanda{client}
}
