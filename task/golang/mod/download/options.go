package download

type Option func(*Options)

type Options struct {
	Flags   []string
	Modules []string
	name    string
}

func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: "go/mod/download",
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// WithFlags sets additional flags to pass to the Go `mod download` command.
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}

// WithModules sets modules to pass to the Go `mod download` command.
func WithModules(modules ...string) Option {
	return func(o *Options) {
		o.Modules = append(o.Flags, modules...)
	}
}
