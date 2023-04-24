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

type getIDOptionImpl struct {
	track     string
	has_track bool
}

func (g *getIDOptionImpl) Track() string  { return g.track }
func (g *getIDOptionImpl) HasTrack() bool { return g.has_track }

type GetIDParams struct {
	Track string `json:"track"`
}

func (o GetIDParams) Options() []GetIDOption {
	return []GetIDOption{
		GetIDTrack(o.Track),
	}
}

// ToBaseOptions converts GetIDOption to an array of BaseOption
func (o *getIDOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{}
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
