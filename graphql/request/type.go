package request

import "context"

type OperationType string

const (
	OptionTypeOperationName      OperationType = "operation_name"
	OptionTypeOperationDirective OperationType = "operation_directive"
)

type Options struct {
	Context       context.Context
	OperationType OperationType
	OperationName string
}

func NewOptions(ctx context.Context, opts ...Option) *Options {
	options := &Options{
		Context: ctx,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}
