package gpool

type OptionsFunc func(*Options)

type Options struct {
	// MaxCapacity 缓存池大小
	MaxCapacity int `json:"max_capacity" yaml:"max_capacity"`

	// MinCapacity 最小容量
	MinCapacity int `json:"min_capacity" yaml:"min_capacity"`

	// NonBlock 异步
	NonBlock bool `json:"non_block" yaml:"non_block"`

	// HideUniqueIdentify 隐藏唯一标识
	HideUniqueIdentify bool `json:"hide_unique_identify" yaml:"hide_unique_identify"`
}

const (
	MIN_CAPACITY = 5
	MAX_CAPACITY = 10000
)

func (o *Options) setDefault() {
	if o.MinCapacity < MIN_CAPACITY {
		o.MinCapacity = MIN_CAPACITY
	}

	if o.MinCapacity > MAX_CAPACITY {
		o.MinCapacity = MAX_CAPACITY
	}

	if o.MaxCapacity < o.MinCapacity {
		o.MaxCapacity = MAX_CAPACITY
	}
}

func dflOptions() *Options {
	return &Options{
		MaxCapacity: MAX_CAPACITY,
		MinCapacity: MIN_CAPACITY,
		NonBlock:    false,
	}
}

func WithOptions() *Options {
	return dflOptions()
}

func WithMaxCapacity(maxCapacity int) OptionsFunc {
	return func(o *Options) {
		o.MaxCapacity = maxCapacity
	}
}

func WithMinCapacity(minCapacity int) OptionsFunc {
	return func(o *Options) {
		o.MinCapacity = minCapacity
	}
}

func WithNonBlock(nonBlock bool) OptionsFunc {
	return func(o *Options) {
		o.NonBlock = nonBlock
	}
}
