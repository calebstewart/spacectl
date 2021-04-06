package client

import "context"

type Client interface {
	// Query executes a single GraphQL query request.
	Query(context.Context, interface{}, map[string]interface{}) error

	// Mutate executes a single GraphQL mutation request.
	Mutate(context.Context, interface{}, map[string]interface{})
}
