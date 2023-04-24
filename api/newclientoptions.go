// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type NewClientOption struct {
	f func(*newClientOptionImpl)
	s string
}

func (o NewClientOption) String() string { return o.s }

type NewClientOptions interface {
	Logger() ApiLogger
	HasLogger() bool
}

func NewClientLogger(logger ApiLogger) NewClientOption {
	return NewClientOption{func(opts *newClientOptionImpl) {
		opts.has_logger = true
		opts.logger = logger
	}, fmt.Sprintf("api.NewClientLogger(ApiLogger %+v)", logger)}
}
func NewClientLoggerFlag(logger *ApiLogger) NewClientOption {
	return NewClientOption{func(opts *newClientOptionImpl) {
		if logger == nil {
			return
		}
		opts.has_logger = true
		opts.logger = *logger
	}, fmt.Sprintf("api.NewClientLogger(ApiLogger %+v)", logger)}
}

type newClientOptionImpl struct {
	logger     ApiLogger
	has_logger bool
}

func (n *newClientOptionImpl) Logger() ApiLogger { return n.logger }
func (n *newClientOptionImpl) HasLogger() bool   { return n.has_logger }

func makeNewClientOptionImpl(opts ...NewClientOption) *newClientOptionImpl {
	res := &newClientOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeNewClientOptions(opts ...NewClientOption) NewClientOptions {
	return makeNewClientOptionImpl(opts...)
}
