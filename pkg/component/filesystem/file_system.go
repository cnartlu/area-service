package filesystem

import "context"

type AwsInterface interface{}

type Uploader interface {
	Upload(ctx context.Context, options ...Option) error
}
