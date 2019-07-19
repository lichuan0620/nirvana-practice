package handlers

import (
	"context"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/service"
	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

func init() {
	cacheService = service.NewCacheService()
}

var cacheService service.CacheService

func CreateCache(ctx context.Context, input *api.Cache) (*api.Cache, error) {
	return cacheService.Create(input)
}

func ListCaches(ctx context.Context, prefix string) ([]api.Cache, error) {
	return cacheService.List(prefix)
}

func GetCache(ctx context.Context, key string) (*api.Cache, error) {
	return cacheService.Get(key)
}

func UpdateCache(ctx context.Context, key string, new *api.Cache) (*api.Cache, error) {
	if new == nil || len(new.Key) == 0 || new.Key != key {
		return nil, errors.ErrorValidationFailed.Error("")
	}

	return cacheService.Update(key, new)
}

func DeleteCache(ctx context.Context, key string) error {
	return cacheService.Delete(key)
}
