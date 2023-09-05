package download

import taskGo "github.com/svengreb/wand/pkg/task/golang"

type Option func(*Options)

type Options struct {
	*taskGo.Options

	Flags      []string
	Inputs     []string
	name       string
	taskGoOpts []taskGo.Option
}

func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: "go/mod/generate",
	}
	for _, o := range opts {
		o(opt)
	}
	opt.Options = taskGo.NewOptions(opt.taskGoOpts...)
	return opt
}

// WithFlags sets additional flags to pass to the Go `generate` command.
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithGoOptions sets shared Go toolchain task options.
func WithGoOptions(goOpts ...taskGo.Option) Option {
	return func(o *Options) {
		o.taskGoOpts = append(o.taskGoOpts, goOpts...)
	}
}

// WithInputs sets inputs to pass to the Go `generate` command.
func WithInputs(inputs ...string) Option {
	return func(o *Options) {
		o.Inputs = append(o.Flags, inputs...)
	}
}
