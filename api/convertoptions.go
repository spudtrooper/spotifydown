// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type ConvertOption struct {
	f func(*convertOptionImpl)
	s string
}

func (o ConvertOption) String() string { return o.s }

type ConvertOptions interface {
	Track() string
	HasTrack() bool
	Verbose() bool
	HasVerbose() bool
	ToBaseOptions() []BaseOption
}

func ConvertTrack(track string) ConvertOption {
	return ConvertOption{func(opts *convertOptionImpl) {
		opts.has_track = true
		opts.track = track
	}, fmt.Sprintf("api.ConvertTrack(string %+v)", track)}
}
func ConvertTrackFlag(track *string) ConvertOption {
	return ConvertOption{func(opts *convertOptionImpl) {
		if track == nil {
			return
		}
		opts.has_track = true
		opts.track = *track
	}, fmt.Sprintf("api.ConvertTrack(string %+v)", track)}
}

func ConvertVerbose(verbose bool) ConvertOption {
	return ConvertOption{func(opts *convertOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.ConvertVerbose(bool %+v)", verbose)}
}
func ConvertVerboseFlag(verbose *bool) ConvertOption {
	return ConvertOption{func(opts *convertOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.ConvertVerbose(bool %+v)", verbose)}
}

type convertOptionImpl struct {
	track       string
	has_track   bool
	verbose     bool
	has_verbose bool
}

func (c *convertOptionImpl) Track() string    { return c.track }
func (c *convertOptionImpl) HasTrack() bool   { return c.has_track }
func (c *convertOptionImpl) Verbose() bool    { return c.verbose }
func (c *convertOptionImpl) HasVerbose() bool { return c.has_verbose }

type ConvertParams struct {
	Track   string `json:"track"`
	Verbose bool   `json:"verbose"`
}

func (o ConvertParams) Options() []ConvertOption {
	return []ConvertOption{
		ConvertTrack(o.Track),
		ConvertVerbose(o.Verbose),
	}
}

// ToBaseOptions converts ConvertOption to an array of BaseOption
func (o *convertOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{}
}

func makeConvertOptionImpl(opts ...ConvertOption) *convertOptionImpl {
	res := &convertOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeConvertOptions(opts ...ConvertOption) ConvertOptions {
	return makeConvertOptionImpl(opts...)
}
