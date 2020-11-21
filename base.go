package gpool

import (
	"context"
	"sync"
	"sync/atomic"
)

type base struct {
	ctx    context.Context
	cancel context.CancelFunc

	running int32
	opt     *Options

	size int

	wg sync.WaitGroup

	log Logger

	ft  fnType
	fn1 ExecArg
	fn2 ExecArgResult
}

type fnType int8

const (
	_ = iota
	fnTypeFunc
	fnTypeArg
	fnTypeArgResult
)

func newBase(ctx context.Context, opt *Options, size int) *base {
	c, cancel := context.WithCancel(ctx)

	return &base{
		ctx:    c,
		cancel: cancel,
		opt:    opt,
		size:   size,
		log:    dfl,
	}
}

func (b *base) Stop() { b.cancel() }
func (b *base) Wait() { b.wg.Wait() }

func (b *base) IsRunning() bool {
	return !atomic.CompareAndSwapInt32(&b.running, 0, 1)
}
