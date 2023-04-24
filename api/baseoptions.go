// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type BaseOption struct {
	f func(*baseOptionImpl)
	s string
}

func (o BaseOption) String() string { return o.s }

type BaseOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func BaseVerbose(verbose bool) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.BaseVerbose(bool %+v)", verbose)}
}
func BaseVerboseFlag(verbose *bool) BaseOption {
	return BaseOption{func(opts *baseOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.BaseVerbose(bool %+v)", verbose)}
}

type baseOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (b *baseOptionImpl) Verbose() bool    { return b.verbose }
func (b *baseOptionImpl) HasVerbose() bool { return b.has_verbose }

func makeBaseOptionImpl(opts ...BaseOption) *baseOptionImpl {
	res := &baseOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeBaseOptions(opts ...BaseOption) BaseOptions {
	return makeBaseOptionImpl(opts...)
}
