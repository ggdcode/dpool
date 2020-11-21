package gpool_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ggdcode/gpool"
)

func TestPoolArg(t *testing.T) {
	p := gpool.New(gpool.WithOptions())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var val int32
	fn := func(ctx context.Context, v interface{}) {
		select {
		case <-ctx.Done():

		default:
			atomic.AddInt32(&val, v.(int32))
		}
	}

	pl := p.CallArg(ctx, fn, 5)
	for i := 0; i < 1000000; i++ {
		err := pl.Submit(int32(1))
		if err != nil {
			atomic.AddInt32(&val, 1)
		}
	}

	p.Wait()
	t.Log("val =", val)
}
