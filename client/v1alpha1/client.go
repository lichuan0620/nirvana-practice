package v1alpha1

import (
	context "context"

	rest "github.com/caicloud/nirvana/rest"
)

// Interface describes v1alpha1 client.
type Interface interface {
	CreateCache(ctx context.Context, input *Cache) (created *Cache, err error)
	ListCaches(ctx context.Context, prefix string) (caches []Cache, err error)
	GetCache(ctx context.Context, key string) (cache *Cache, err error)
	UpdateCache(ctx context.Context, key string, input *Cache) (updated *Cache, err error)
	DeleteCache(ctx context.Context, key string) (err error)
}

// Client for version v1.
type Client struct {
	rest *rest.Client
}

// NewClient creates a new client.
func NewClient(cfg *rest.Config) (*Client, error) {
	client, err := rest.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// MustNewClient creates a new client or panic if an error occurs.
func MustNewClient(cfg *rest.Config) *Client {
	client, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *Client) CreateCache(ctx context.Context, input *Cache) (created *Cache, err error) {
	created = new(Cache)
	err = c.rest.Request("POST", 201, "/api/v1alpha1/caches").
		Body("application/json", input).
		Data(created).
		Do(ctx)

	return
}

func (c *Client) ListCaches(ctx context.Context, prefix string) (caches []Cache, err error) {
	err = c.rest.Request("GET", 200, "/api/v1alpha1/caches").
		Query("prefix", prefix).
		Data(&caches).
		Do(ctx)
	return
}

func (c *Client) GetCache(ctx context.Context, key string) (cache *Cache, err error) {
	cache = new(Cache)
	err = c.rest.Request("GET", 200, "/api/v1alpha1/caches/{key}").
		Path("key", key).
		Data(cache).
		Do(ctx)
	return
}

func (c *Client) UpdateCache(ctx context.Context, key string, input *Cache) (updated *Cache, err error) {
	updated = new(Cache)
	err = c.rest.Request("PUT", 200, "/api/v1alpha1/caches/{key}").
		Path("key", key).
		Body("application/json", input).
		Data(updated).
		Do(ctx)
	return
}

func (c *Client) DeleteCache(ctx context.Context, key string) (err error) {
	err = c.rest.Request("DELETE", 204, "/api/v1alpha1/caches/{key}").
		Path("key", key).
		Do(ctx)
	return
}
