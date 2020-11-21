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
		p.log.Warn("WAS_RUN")
		return p
	}

	p.wg.Add(p.size)
	p.log.Debug("IS_RUNNING")

	for i := 0; i < p.size; i++ {
		i := i
		go func() {
			p.fn(p.ctx)
			p.log.WithField("index", i).Debug("WAS_STOPPED")
			p.wg.Done()
		}()
	}

	return p
}
