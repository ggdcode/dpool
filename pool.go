package gpool

import (
	"context"
)

type pool struct {
	*base

	fn ExecFunc
}

func newPool(ctx context.Context, opt *Options, size int, fn ExecFunc) *pool {
	return &pool{
		base: newBase(ctx, opt, size),
		fn:   fn,
	}
}

func (p *pool) thread() *pool {
	if p.IsRunning() {
		p.log.Warn("the goroutine pool was run!")
		return p
	}

	p.wg.Add(p.size)
	p.log.Debug("the goroutine pool is running!")

	for i := 0; i < p.size; i++ {
		i := i
		go func() {
			p.fn(p.ctx)
			p.wg.Done()
			p.log.WithField("index", i).Debug("the goroutine was stopped")
		}()
	}

	return p
}
