package gpool

import (
	"context"
	"errors"
)

type WorkerFunc func(ctx context.Context, wm WorkerManager, fn ExecArgResult) Worker

type Worker interface {
	OnMessage(ctx context.Context, param interface{}) (interface{}, error)
}

type worker struct {
	wm WorkerManager

	fn ExecArgResult

	paramCh   chan interface{}
	onMessage func(ctx context.Context, param interface{}) (interface{}, error)
}

func newWorker(ctx context.Context, wm WorkerManager, fn ExecArgResult) Worker {
	w := &worker{
		wm:      wm,
		fn:      fn,
		paramCh: make(chan interface{}, 1),
	}

	if wm.Async() {
		w.wm.Pool().WgAdd(1)
		w.onMessage = w.onMessageAsync
		go w.run(ctx)
	} else {
		w.onMessage = w.OnMessageSync
	}

	w.wm.AddWorker(ctx, w)

	return w
}

func (w *worker) run(ctx context.Context) {
	defer w.wm.Pool().WgDone()

	for {
		select {
		case <-ctx.Done():
			return

		case param := <-w.paramCh:
			w.fn(ctx, param)
		}
	}
}

func (w *worker) OnMessage(ctx context.Context, param interface{}) (interface{}, error) {
	return w.onMessage(ctx, param)
}

func (w *worker) onMessageAsync(ctx context.Context, param interface{}) (interface{}, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("STOPPED")
	case w.paramCh <- param:
	}

	w.wm.AddWorker(ctx, w)

	return nil, nil
}

func (w *worker) OnMessageSync(ctx context.Context, param interface{}) (interface{}, error) {
	res := w.fn(ctx, param)
	w.wm.AddWorker(ctx, w)

	return res, nil
}
