package middlewares

import (
	"context"

	"github.com/caicloud/nirvana/definition"

	"github.com/lichuan0620/nirvana-practice/client"
)

const (
	keyCache int = iota
)

func WithClientInterface(client client.Interface) definition.Middleware {
	return func(ctx context.Context, chain definition.Chain) error {
		ctx = context.WithValue(ctx, keyCache, client)
		return chain.Continue(ctx)
	}
}

func GetClientInterface(ctx context.Context) (client.Interface, bool) {
	client, ok := ctx.Value(keyCache).(client.Interface)
	return client, ok
}
