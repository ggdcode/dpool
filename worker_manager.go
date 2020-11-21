package gpool

import (
	"context"
	"errors"
)

type WorkerManager interface {
	Pool() Pool
	Async() bool
	AddWorker(ctx context.Context, w Worker)
	OnMessage(ctx context.Context, param interface{}) (interface{}, error)
}

type workerMgr struct {
	pool Pool

	async bool
	cap   int
	size  int
	auto  bool

	fn ExecArgResult

	paramCh chan interface{} // 参数数据
	workers chan Worker
}

type WorkerManagerFunc func(ctx context.Context, pool Pool, opt *WorkerManagerOptions, fn ExecArgResult) WorkerManager

func newWorkerMgr(ctx context.Context, pool Pool, opt *WorkerManagerOptions, fn ExecArgResult) WorkerManager {
	wm := &workerMgr{
		pool: pool,

		async: opt.Async,
		cap:   opt.Cap,
		size:  opt.Size,
		auto:  opt.Expansion,

		fn: fn,

		paramCh: make(chan interface{}, opt.Cap),
		workers: make(chan Worker, opt.Size),
	}

	if wm.async {
		wm.pool.WgAdd(1)
		go wm.run(ctx)
	}

	wm.dynamicCap(ctx, opt.Size)
	return wm
}

func (wm *workerMgr) run(ctx context.Context) {
	defer wm.pool.WgDone()
	wm.pool.GetLogger().Debug("IS_RUNNING")

	for {
		select {
		case <-ctx.Done():
			return

		case param := <-wm.paramCh:
			select {
			case <-ctx.Done():
				return

			case w := <-wm.workers:
				w.OnMessage(ctx, param)
			}
		}
	}
}

func (wm *workerMgr) AddWorker(ctx context.Context, w Worker) {
	select {
	case <-ctx.Done():
	case wm.workers <- w:
	}
}

func (wm *workerMgr) OnMessage(ctx context.Context, param interface{}) (interface{}, error) {
	if wm.async {
		select {
		case <-ctx.Done():
		case wm.paramCh <- param:
		}

		return nil, nil
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("STOPPED")
	case w := <-wm.workers:
		return w.OnMessage(ctx, param)
	}
}

// dynamicCap
//! TOTO
func (wm *workerMgr) dynamicCap(ctx context.Context, size int) {
	for i := 0; i < size; i++ {
		wm.pool.Options().FnNewWorker(ctx, wm, wm.fn)
	}
}

func (wm *workerMgr) Pool() Pool  { return wm.pool }
func (wm *workerMgr) Async() bool { return wm.async }
