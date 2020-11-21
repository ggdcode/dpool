package gpool_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ggdcode/gpool"
)

func TestPool(t *testing.T) {
	p := gpool.New(gpool.WithOptions())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var val int32
	fn := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			atomic.AddInt32(&val, 1)
		}
	}

	p.Call(ctx, fn, 10)
	p.Wait()

	t.Log("val =", val)
}
