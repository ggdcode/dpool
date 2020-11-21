package gpool

import "github.com/ggdcode/glog"

type Logger = glog.Logger

var dfl = glog.GetDefaultLog()

func (b *base) SetLogger(l glog.Logger) {
	b.log = l
}

func (b *base) GetLogger() glog.Logger {
	return b.log
}
