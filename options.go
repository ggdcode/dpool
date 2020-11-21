package gpool

type OptionsFunc func(*Options)

type Options struct {
	// Expansion
	Expansion bool `json:"expansion" yaml:"expansion"`

	// MaxCapacity 缓存池大小
	MaxCapacity int `json:"max_capacity" yaml:"max_capacity"`

	// NonBlock 异步
	NonBlock bool `json:"non_block" yaml:"non_block"`

	// HideUniqueIdentify 隐藏唯一标识
	HideUniqueIdentify bool `json:"hide_unique_identify" yaml:"hide_unique_identify"`

	FnNewWorkerManager WorkerManagerFunc
	FnNewWorker        WorkerFunc
}

const (
	MAX_CAPACITY = 100000
	MIN_CAPACITY = 5
)

func (o *Options) setDefault() {
	if o.MaxCapacity > MAX_CAPACITY {
		o.MaxCapacity = MAX_CAPACITY
	} else if o.MaxCapacity < MIN_CAPACITY {
		o.MaxCapacity = MIN_CAPACITY
	}

	if o.FnNewWorkerManager == nil {
		o.FnNewWorkerManager = newWorkerMgr
	}

	if o.FnNewWorker == nil {
		o.FnNewWorker = newWorker
	}
}

func dflOptions() *Options {
	return &Options{
		MaxCapacity: MAX_CAPACITY,
		NonBlock:    false,
	}
}

func WithOptions() *Options {
	return dflOptions()
}

func WithMaxCapacity(maxCapacity int) OptionsFunc {
	return func(o *Options) {
		if maxCapacity <= MAX_CAPACITY && maxCapacity >= MIN_CAPACITY {
			o.MaxCapacity = maxCapacity
		}
	}
}

func WithNonBlock(nonBlock bool) OptionsFunc {
	return func(o *Options) {
		o.NonBlock = nonBlock
	}
}
