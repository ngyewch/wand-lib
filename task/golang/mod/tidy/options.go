package download

type Option func(*Options)

type Options struct {
	Flags []string
	name  string
}

func NewOptions(opts ...Option) *Options {
	opt := &Options{
		name: "go/mod/tidy",
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

// WithFlags sets additional flags to pass to the Go `mod tidy` command.
func WithFlags(flags ...string) Option {
	return func(o *Options) {
		o.Flags = append(o.Flags, flags...)
	}
}
