// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import "fmt"

type ProgressOption struct {
	f func(*progressOptionImpl)
	s string
}

func (o ProgressOption) String() string { return o.s }

type ProgressOptions interface {
	TaskID() string
	HasTaskID() bool
	ToBaseOptions() []BaseOption
}

func ProgressTaskID(taskID string) ProgressOption {
	return ProgressOption{func(opts *progressOptionImpl) {
		opts.has_taskID = true
		opts.taskID = taskID
	}, fmt.Sprintf("api.ProgressTaskID(string %+v)", taskID)}
}
func ProgressTaskIDFlag(taskID *string) ProgressOption {
	return ProgressOption{func(opts *progressOptionImpl) {
		if taskID == nil {
			return
		}
		opts.has_taskID = true
		opts.taskID = *taskID
	}, fmt.Sprintf("api.ProgressTaskID(string %+v)", taskID)}
}

type progressOptionImpl struct {
	taskID     string
	has_taskID bool
}

func (p *progressOptionImpl) TaskID() string  { return p.taskID }
func (p *progressOptionImpl) HasTaskID() bool { return p.has_taskID }

type ProgressParams struct {
	TaskID string `json:"task_id"`
}

func (o ProgressParams) Options() []ProgressOption {
	return []ProgressOption{
		ProgressTaskID(o.TaskID),
	}
}

// ToBaseOptions converts ProgressOption to an array of BaseOption
func (o *progressOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{}
}

func makeProgressOptionImpl(opts ...ProgressOption) *progressOptionImpl {
	res := &progressOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeProgressOptions(opts ...ProgressOption) ProgressOptions {
	return makeProgressOptionImpl(opts...)
}
