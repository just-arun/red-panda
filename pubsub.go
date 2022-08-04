package redpanda

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

func (st *redPanda) Publish(eventName string, payload []byte) (err error) {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Topic: eventName, Value: payload}
	st.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
	return
}

func (st *redPanda) Subscribe(ctx context.Context, event string, cb func(err error, msg *kgo.Record)) {
	for {
		fetches := st.PollFetches(ctx)
		fetches.EachPartition(func(p kgo.FetchTopicPartition) {
			p.EachRecord(func(record *kgo.Record) {
				if record.Topic == event {
					cb(p.Err, record)
				}
			})
		})
	}
}
