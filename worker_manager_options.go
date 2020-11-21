package gpool

type WorkerManagerOptions struct {
	// Async 异步处理
	Async bool

	// Cap 消息最大容量
	Cap int

	// 工作协程数
	Size int

	// 允许自动扩容
	Expansion bool
}
