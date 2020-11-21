package gpool

type WorkerOptions struct {
	// Cap 消息最大容量
	Cap int

	// 工作协程数
	Size int

	// 允许自动扩容
	Auto bool
}

func WithWorkerOptions() *WorkerOptions {
	return &WorkerOptions{
		Cap:  MIN_CAPACITY,
		Size: MIN_CAPACITY,
		Auto: false,
	}
}
