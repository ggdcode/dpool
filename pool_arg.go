package gpool

import "context"

type poolArg struct {
	*base

	wmOpt *WorkerManagerOptions
	wm    WorkerManager
	fn    ExecArgResult
}

func newPoolArg(ctx context.Context, opt *Options, size int, fn ExecArg) *poolArg {
	return &poolArg{
		base: newBase(ctx, opt, size),
		wmOpt: &WorkerManagerOptions{
			Async:     true,
			Cap:       opt.MaxCapacity,
			Size:      size,
			Expansion: opt.Expansion,
		},
		fn: func(ctx context.Context, val interface{}) interface{} {
			fn(ctx, val)
			return nil
		},
	}
}

func (p *poolArg) thread() *poolArg {
	p.wm = p.opt.FnNewWorkerManager(p.ctx, p, p.wmOpt, p.fn)
	return p
}

func (p *poolArg) Submit(param interface{}) error {
	_, err := p.wm.OnMessage(p.ctx, param)
	return err
}
