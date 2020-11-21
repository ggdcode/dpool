package gpool

import (
	"context"

	"github.com/ggdcode/glog"
)

type (
	ExecFunc      func(context.Context)
	ExecArg       func(context.Context, interface{})
	ExecArgResult func(context.Context, interface{}) interface{}
)

type WaiterHandler interface {
	Stop()
	Wait()
}

type LoggerHandler interface {
	SetLogger(l glog.Logger)
	GetLogger() glog.Logger
}

// PoolGroup 协程池接口
type PoolGroup interface {
	WaiterHandler

	Call(context.Context, func(context.Context), ...int) PoolFunc

	CallArg(context.Context, func(context.Context, interface{}), ...int) PoolArg

	CallArgResult(context.Context, func(context.Context, interface{}) interface{}, ...int) PoolArg
}
type Pool interface {
	Options() *Options

	WgAdd(n int)
	WgDone()

	WaiterHandler
	LoggerHandler
}

type PoolFunc interface {
	Pool
}

type PoolArg interface {
	Pool

	Submit(param interface{}) error
}

type PoolArgResult interface {
	Pool
}
