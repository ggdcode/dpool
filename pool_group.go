package gpool

import (
	"context"
	"math/rand"
	"sync/atomic"
)

var poolN uint32 = 0x10101010

type poolGroup struct {
	opt *Options

	w []WaiterHandler
}

func New(opt *Options) PoolGroup {
	opt.setDefault()

	return &poolGroup{opt: opt}
}

func (p *poolGroup) addWaiter(w WaiterHandler) { p.w = append(p.w, w) }

func (p *poolGroup) Call(ctx context.Context, fn func(context.Context), size ...int) PoolFunc {
	var siz int

	if len(size) > 0 && size[0] > 0 {
		siz = size[0]
	} else {
		siz = p.opt.MinCapacity
	}

	b := newPool(ctx, p.opt, siz, fn)
	p.addWaiter(b)

	if !p.opt.HideUniqueIdentify {
		b.log = b.log.WithField("pool_id", getPoolCount())
	}

	return b.thread()
}

func (p *poolGroup) CallArg(ctx context.Context, fn func(context.Context, interface{}), size ...int) PoolArg {
	return nil
}

func (p *poolGroup) CallArgResult(ctx context.Context, fn func(context.Context, interface{}) interface{}, size ...int) PoolArg {
	return nil
}

func (p *poolGroup) Stop() {
	for _, w := range p.w {
		w.Stop()
	}
}

func (p *poolGroup) Wait() {
	for _, w := range p.w {
		w.Wait()
	}
}

func addPoolCount() {
	atomic.AddUint32(&poolN, (1+uint32(rand.Int31()))%100)
}

func getPoolCount() uint32 {
	addPoolCount()

	return atomic.LoadUint32(&poolN)
}
