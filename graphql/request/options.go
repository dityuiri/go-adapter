package request

import "context"

type Option func(options *Options)

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func WithOperationName(opName string) Option {
	return func(o *Options) {
		o.OperationName = opName
	}
}
