// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type MetadataOption struct {
	f func(*metadataOptionImpl)
	s string
}

func (o MetadataOption) String() string { return o.s }

type MetadataOptions interface {
	Track() string
	HasTrack() bool
	Verbose() bool
	HasVerbose() bool
	ToBaseOptions() []BaseOption
}

func MetadataTrack(track string) MetadataOption {
	return MetadataOption{func(opts *metadataOptionImpl) {
		opts.has_track = true
		opts.track = track
	}, fmt.Sprintf("api.MetadataTrack(string %+v)", track)}
}
func MetadataTrackFlag(track *string) MetadataOption {
	return MetadataOption{func(opts *metadataOptionImpl) {
		if track == nil {
			return
		}
		opts.has_track = true
		opts.track = *track
	}, fmt.Sprintf("api.MetadataTrack(string %+v)", track)}
}

func MetadataVerbose(verbose bool) MetadataOption {
	return MetadataOption{func(opts *metadataOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("api.MetadataVerbose(bool %+v)", verbose)}
}
func MetadataVerboseFlag(verbose *bool) MetadataOption {
	return MetadataOption{func(opts *metadataOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("api.MetadataVerbose(bool %+v)", verbose)}
}

type metadataOptionImpl struct {
	track       string
	has_track   bool
	verbose     bool
	has_verbose bool
}

func (m *metadataOptionImpl) Track() string    { return m.track }
func (m *metadataOptionImpl) HasTrack() bool   { return m.has_track }
func (m *metadataOptionImpl) Verbose() bool    { return m.verbose }
func (m *metadataOptionImpl) HasVerbose() bool { return m.has_verbose }

type MetadataParams struct {
	Track   string `json:"track"`
	Verbose bool   `json:"verbose"`
}

func (o MetadataParams) Options() []MetadataOption {
	return []MetadataOption{
		MetadataTrack(o.Track),
		MetadataVerbose(o.Verbose),
	}
}

// ToBaseOptions converts MetadataOption to an array of BaseOption
func (o *metadataOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseVerbose(o.Verbose()),
	}
}

func makeMetadataOptionImpl(opts ...MetadataOption) *metadataOptionImpl {
	res := &metadataOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeMetadataOptions(opts ...MetadataOption) MetadataOptions {
	return makeMetadataOptionImpl(opts...)
}
