// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type BaseOption struct {
	f func(*baseOptionImpl)
	s string
}

func (o BaseOption) String() string { return o.s }

type BaseOptions interface {
}

type baseOptionImpl struct {
}

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
