package graphql

import (
	"context"
	"net/http"

	"github.com/dityuiri/go-adapter/graphql/request"

	graphQlClient "github.com/hasura/go-graphql-client"
)

//go:generate mockgen -destination=mock/client.go -package=mock . IClient

type Client struct {
	Context         context.Context
	Configuration   *Configuration
	doer            Doer
	debug           bool
	requestModifier graphQlClient.RequestModifier
	client          *graphQlClient.Client
}

type IClient interface {
	Query(query interface{}, variables map[string]interface{}, opts ...request.Option) error
	Mutate(mutation interface{}, variables map[string]interface{}, opts ...request.Option) error
	QueryRaw(query interface{}, variables map[string]interface{}, opts ...request.Option) ([]byte, error)
	MutateRaw(mutation interface{}, variables map[string]interface{}, opts ...request.Option) ([]byte, error)
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
