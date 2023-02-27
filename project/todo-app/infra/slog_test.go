package infra

import (
	"context"
	"testing"
)

func TestInfoMsg(t *testing.T) {
	Info().Str("str", "A").Int("int", 1).Any("any", struct{}{}).Msg(context.Background(), "message")
}

func TestInfoSend(t *testing.T) {
	Info().Str("str", "A").Int("int", 1).Any("any", struct{}{}).Send(context.Background())
}

func TestInfoSendParallel(t *testing.T) {
	l := Info()

	c := make(chan struct{}, 5)
	for i := 0; i < 5; i++ {
		go func(ni int) {
			l.Int("count", ni)
			c <- struct{}{}
		}(i + 1)
	}
	for i := 0; i < 5; i++ {
		_ = <-c
	}
	l.Send(context.Background())
}
