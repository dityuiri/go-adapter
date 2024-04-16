package graphql

import (
	"context"

	"github.com/dityuiri/go-adapter/graphql/request"
	graphQlClient "github.com/hasura/go-graphql-client"
)

const (
	defaultDebug = false
)

func NewClient(ctx context.Context, config *Configuration, opts ...Option) IClient {
	client := &Client{
		Context:       ctx,
		Configuration: config,
		debug:         defaultDebug,
	}

	client = client.set(opts...)
	client.client = graphQlClient.NewClient(client.Configuration.ClientURL, client.doer).WithDebug(client.debug)

	if client.requestModifier != nil {
		client.client.WithRequestModifier(client.requestModifier)
	}

	return client
}

func (c *Client) set(opts ...Option) *Client {
	for _, fn := range opts {
		fn(c)
	}
	return c
}

func (c *Client) Query(query interface{}, variables map[string]interface{}, opts ...request.Option) error {
	options := request.NewOptions(c.Context, opts...)
	return c.client.Query(options.Context, query, variables)
}

func (c *Client) Mutate(mutation interface{}, variables map[string]interface{}, opts ...request.Option) error {
	options := request.NewOptions(c.Context, opts...)
	return c.client.Mutate(options.Context, mutation, variables)
}

func (c *Client) QueryRaw(query interface{}, variables map[string]interface{}, opts ...request.Option) ([]byte, error) {
	options := request.NewOptions(c.Context, opts...)
	if options.OperationName != "" {
		return c.client.NamedQueryRaw(options.Context, options.OperationName, query, variables)
	}

	return c.client.QueryRaw(options.Context, query, variables)
}

func (c *Client) MutateRaw(mutation interface{}, variables map[string]interface{}, opts ...request.Option) ([]byte, error) {
	options := request.NewOptions(c.Context, opts...)
	if options.OperationName != "" {
		return c.client.NamedMutateRaw(options.Context, options.OperationName, mutation, variables)
	}

	return c.client.MutateRaw(options.Context, mutation, variables)
}
