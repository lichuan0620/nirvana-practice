package middlewares

import (
	"context"

	"github.com/caicloud/nirvana/definition"

	"github.com/lichuan0620/nirvana-practice/client"
	"github.com/lichuan0620/nirvana-practice/pkg/apis/service"
)

const (
	keyCache int = iota
	keyCacheService
	KeyProductService
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

func WithCacheService(cacheService service.CacheService) definition.Middleware {
	return func(ctx context.Context, chain definition.Chain) error {
		ctx = context.WithValue(ctx, keyCacheService, cacheService)
		return chain.Continue(ctx)
	}
}

func GetCacheService(ctx context.Context) (service.CacheService, bool) {
	cacheService, ok := ctx.Value(keyCacheService).(service.CacheService)
	return cacheService, ok
}

func WithProductService(productService service.ProductService) definition.Middleware {
	return func(ctx context.Context, chain definition.Chain) error {
		ctx = context.WithValue(ctx, KeyProductService, productService)
		return chain.Continue(ctx)
	}
}

func GetProductService(ctx context.Context) (service.ProductService, bool) {
	productService, ok := ctx.Value(KeyProductService).(service.ProductService)
	return productService, ok
}
