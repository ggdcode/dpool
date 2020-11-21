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

// Pool 协程池接口
type Pool interface {
	WaiterHandler

	Call(context.Context, func(context.Context), ...int) PoolFunc

	CallArg(context.Context, func(context.Context, interface{}), ...int) PoolArg

	CallArgResult(context.Context, func(context.Context, interface{}) interface{}, ...int) PoolArg
}

type PoolFunc interface {
	WaiterHandler

	LoggerHandler
}

type PoolArg interface {
	WaiterHandler

	LoggerHandler
}

type PoolArgResult interface {
	WaiterHandler

	LoggerHandler
}
