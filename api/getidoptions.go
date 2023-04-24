// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type GetIDOption struct {
	f func(*getIDOptionImpl)
	s string
}

func (o GetIDOption) String() string { return o.s }

type GetIDOptions interface {
	Track() string
	HasTrack() bool
	Verbose() bool
	HasVerbose() bool
	ToBaseOptions() []BaseOption
}

func GetIDTrack(track string) GetIDOption {
	return GetIDOption{func(opts *getIDOptionImpl) {
		opts.has_track = true
		opts.track = track
	}, fmt.Sprintf("api.GetIDTrack(string %+v)", track)}
}
func GetIDTrackFlag(track *string) GetIDOption {
	return GetIDOption{func(opts *getIDOptionImpl) {
		if track == nil {
			return
		}
		opts.has_track = true
		opts.track = *track
	}, fmt.Sprintf("api.GetIDTrack(string %+v)", track)}
}

func GetIDVerbose(verbose bool) GetIDOption {
	return GetIDOption{func(opts *getIDOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.GetIDVerbose(bool %+v)", verbose)}
}
func GetIDVerboseFlag(verbose *bool) GetIDOption {
	return GetIDOption{func(opts *getIDOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.GetIDVerbose(bool %+v)", verbose)}
}

type getIDOptionImpl struct {
	track       string
	has_track   bool
	verbose     bool
	has_verbose bool
}

func (g *getIDOptionImpl) Track() string    { return g.track }
func (g *getIDOptionImpl) HasTrack() bool   { return g.has_track }
func (g *getIDOptionImpl) Verbose() bool    { return g.verbose }
func (g *getIDOptionImpl) HasVerbose() bool { return g.has_verbose }

type GetIDParams struct {
	Track   string `json:"track"`
	Verbose bool   `json:"verbose"`
}

func (o GetIDParams) Options() []GetIDOption {
	return []GetIDOption{
		GetIDTrack(o.Track),
		GetIDVerbose(o.Verbose),
	}
}

// ToBaseOptions converts GetIDOption to an array of BaseOption
func (o *getIDOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseVerbose(o.Verbose()),
	}
}

func makeGetIDOptionImpl(opts ...GetIDOption) *getIDOptionImpl {
	res := &getIDOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGetIDOptions(opts ...GetIDOption) GetIDOptions {
	return makeGetIDOptionImpl(opts...)
}
