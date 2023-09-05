package download

import (
	"github.com/svengreb/wand/pkg/app"
	"github.com/svengreb/wand/pkg/task"
	taskGo "github.com/svengreb/wand/pkg/task/golang"
)

type Task struct {
	ac   app.Config
	opts *Options
}

func New(ac app.Config, opts ...Option) *Task {
	opt := NewOptions(opts...)
	return &Task{
		ac:   ac,
		opts: opt,
	}
}

func (t *Task) BuildParams() []string {
	params := []string{"generate"}

	if len(t.opts.taskGoOpts) > 0 {
		params = append(params, taskGo.BuildGoOptions(t.opts.taskGoOpts...)...)
	}

	if len(t.opts.Flags) > 0 {
		params = append(params, t.opts.Flags...)
	}

	if len(t.opts.Inputs) > 0 {
		params = append(params, t.opts.Inputs...)
	} else {
		params = append(params, "./...")
	}

	return params
}

func (t *Task) Env() map[string]string {
	return t.opts.Env
}

func (t *Task) Kind() task.Kind {
	return task.KindExec
}

func (t *Task) Name() string {
	return t.opts.name
}

func (t *Task) Options() task.Options {
	return *t.opts
}
