package gpool

import (
	"context"
	"math/rand"
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
}

func newBase(ctx context.Context, opt *Options, size int) *base {
	c, cancel := context.WithCancel(ctx)
	addPoolN()

	return &base{
		ctx:    c,
		cancel: cancel,
		opt:    opt,
		size:   size,
		log:    dfl,
	}
}

func (b *base) Options() *Options { return b.opt }
func (b *base) Stop()             { b.cancel() }
func (b *base) Wait() {
	b.wg.Wait()
	b.log.Debug("POOL_ENDED")
}
func (b *base) WgAdd(n int) { b.wg.Add(n) }
func (b *base) WgDone()     { b.wg.Done() }

func (b *base) IsRunning() bool {
	return !atomic.CompareAndSwapInt32(&b.running, 0, 1)
}

var poolN uint32 = 0x10101010

func addPoolN() {
	atomic.AddUint32(&poolN, (1+uint32(rand.Int31()))%100)
}

func getPoolN() uint32 {
	return atomic.LoadUint32(&poolN)
}
