// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type DownloadOption struct {
	f func(*downloadOptionImpl)
	s string
}

func (o DownloadOption) String() string { return o.s }

type DownloadOptions interface {
	Id() string
	HasId() bool
	Verbose() bool
	HasVerbose() bool
	ToBaseOptions() []BaseOption
}

func DownloadId(id string) DownloadOption {
	return DownloadOption{func(opts *downloadOptionImpl) {
		opts.has_id = true
		opts.id = id
	}, fmt.Sprintf("api.DownloadId(string %+v)", id)}
}
func DownloadIdFlag(id *string) DownloadOption {
	return DownloadOption{func(opts *downloadOptionImpl) {
		if id == nil {
			return
		}
		opts.has_id = true
		opts.id = *id
	}, fmt.Sprintf("api.DownloadId(string %+v)", id)}
}

func DownloadVerbose(verbose bool) DownloadOption {
	return DownloadOption{func(opts *downloadOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.DownloadVerbose(bool %+v)", verbose)}
}
func DownloadVerboseFlag(verbose *bool) DownloadOption {
	return DownloadOption{func(opts *downloadOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.DownloadVerbose(bool %+v)", verbose)}
}

type downloadOptionImpl struct {
	id          string
	has_id      bool
	verbose     bool
	has_verbose bool
}

func (d *downloadOptionImpl) Id() string       { return d.id }
func (d *downloadOptionImpl) HasId() bool      { return d.has_id }
func (d *downloadOptionImpl) Verbose() bool    { return d.verbose }
func (d *downloadOptionImpl) HasVerbose() bool { return d.has_verbose }

type DownloadParams struct {
	Id      string `json:"id"`
	Verbose bool   `json:"verbose"`
}

func (o DownloadParams) Options() []DownloadOption {
	return []DownloadOption{
		DownloadId(o.Id),
		DownloadVerbose(o.Verbose),
	}
}

// ToBaseOptions converts DownloadOption to an array of BaseOption
func (o *downloadOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseVerbose(o.Verbose()),
	}
}

func makeDownloadOptionImpl(opts ...DownloadOption) *downloadOptionImpl {
	res := &downloadOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeDownloadOptions(opts ...DownloadOption) DownloadOptions {
	return makeDownloadOptionImpl(opts...)
}
