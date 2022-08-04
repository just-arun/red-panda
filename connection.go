package redpanda

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewConnection(opts ...kgo.Opt) *kgo.Client {
	cl, err := kgo.NewClient(
		opts...
	)
	if err != nil {
		panic(err)
	}
	return cl
}
