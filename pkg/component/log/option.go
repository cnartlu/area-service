package log

type Option func(*Logger)

func WithName(name string) Option {
	return func(o *Logger) {
		o.name = name
	}
}
