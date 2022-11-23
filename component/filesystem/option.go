package filesystem

import "context"

type Option interface {
	// IsExist(exists bool)
	// WithContext(ctx context.Context)
}

type HandleFunc func(Option)

func IsExist(exists bool) HandleFunc {
	return func(r Option) {
		// r.IsExist(exists)
	}
}

func WithContext(ctx context.Context) HandleFunc {
	return func(r Option) {
		// r.WithContext(ctx)
	}
}
