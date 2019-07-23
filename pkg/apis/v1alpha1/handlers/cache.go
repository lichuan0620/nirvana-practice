package handlers

import (
	"context"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/middlewares"
	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

func CreateCache(ctx context.Context, input *api.Cache) (*api.Cache, error) {
	cacheService, exist := middlewares.GetCacheService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return cacheService.Create(input)
}

func ListCaches(ctx context.Context, prefix string) ([]api.Cache, error) {
	cacheService, exist := middlewares.GetCacheService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return cacheService.List(prefix)
}

func GetCache(ctx context.Context, key string) (*api.Cache, error) {
	cacheService, exist := middlewares.GetCacheService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return cacheService.Get(key)
}

func UpdateCache(ctx context.Context, key string, new *api.Cache) (*api.Cache, error) {
	if new == nil || len(new.Key) == 0 || new.Key != key {
		return nil, errors.ErrorValidationFailed.Error("")
	}

	cacheService, exist := middlewares.GetCacheService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return cacheService.Update(key, new)
}

func DeleteCache(ctx context.Context, key string) error {
	cacheService, exist := middlewares.GetCacheService(ctx)
	if !exist {
		return errors.ErrorInternal.Error("")
	}

	return cacheService.Delete(key)
}
