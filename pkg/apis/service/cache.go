package service

import (
	"strings"
	"sync"

	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

type CacheService interface {
	Create(*api.Cache) (*api.Cache, error)
	List(filterRule string) ([]api.Cache, error)
	Update(key string, new *api.Cache) (*api.Cache, error)
	Delete(key string) error
	Get(key string) (result *api.Cache, err error)
}

func NewCacheService() CacheService {
	return &cacheImpl{
		table: make(map[string]api.Cache),
	}
}

type cacheImpl struct {
	sync.Mutex
	table map[string]api.Cache
}

func (c *cacheImpl) Create(input *api.Cache) (*api.Cache, error) {
	if len(input.Key) == 0 {
		return nil, errors.ErrorInvalidField.Error("key")
	}

	c.Lock()
	defer c.Unlock()

	_, ok := c.table[input.Key]
	if ok {
		return nil, errors.ErrorAlreadyExist.Error()
	}

	c.table[input.Key] = *input

	return input, nil
}

func (c *cacheImpl) List(prefix string) ([]api.Cache, error) {
	c.Lock()
	defer c.Unlock()

	caches := make([]api.Cache, 0, len(c.table))

	for key, row := range c.table {
		if strings.HasPrefix(key, prefix) {
			caches = append(caches, row)
		}
	}

	return caches, nil
}

func (c *cacheImpl) Update(key string, new *api.Cache) (*api.Cache, error) {
	c.Lock()
	defer c.Unlock()

	_, ok := c.table[key]
	if !ok {
		return nil, errors.ErrorNotFound.Error()
	}

	c.table[key] = *new

	return new, nil
}

func (c *cacheImpl) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	_, ok := c.table[key]
	if !ok {
		return errors.ErrorNotFound.Error()
	}

	delete(c.table, key)

	return nil
}

func (c *cacheImpl) Get(key string) (*api.Cache, error) {
	c.Lock()
	defer c.Unlock()

	row, ok := c.table[key]
	if !ok {
		return nil, errors.ErrorNotFound.Error()
	}

	return &row, nil
}
