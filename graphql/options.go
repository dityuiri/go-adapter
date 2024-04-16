package graphql

import (
	"github.com/hasura/go-graphql-client"
)

type Option func(*Client)

func WithDoer(doer Doer) Option {
	return func(c *Client) {
		c.doer = doer
	}
}

func WithDebug(debugOn bool) Option {
	return func(c *Client) {
		c.debug = debugOn
	}
}

func WithRequestModifier(req graphql.RequestModifier) Option {
	return func(c *Client) {
		c.requestModifier = req
	}
}
