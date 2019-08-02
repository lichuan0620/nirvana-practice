package service

import (
	"context"
	"encoding/json"

	"github.com/lichuan0620/nirvana-practice/client"
	v1alpha1 "github.com/lichuan0620/nirvana-practice/client/v1alpha1"
	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

type ProductService interface {
	Create(client.Interface, *api.Product) (*api.Product, error)
	List(client.Interface) ([]api.Product, error)
	Update(cli client.Interface, name string, new *api.Product) (*api.Product, error)
	Delete(cli client.Interface, name string) error
	Get(client client.Interface, name string) (result *api.Product, err error)
}

func NewProductService() ProductService {
	return &productImpl{}
}

type productImpl struct{}

const prefix string = "product-"

func (p *productImpl) Create(cli client.Interface, input *api.Product) (*api.Product, error) {
	cache := new(v1alpha1.Cache)
	cache.Key = prefix + input.Name
	cache.Value = *input

	cached, err := cli.V1Alpha1().CreateCache(context.Background(), cache)
	if err != nil {
		return nil, err
	}

	product, err := convertToProduct(cached.Value)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productImpl) List(cli client.Interface) ([]api.Product, error) {
	cached, err := cli.V1Alpha1().ListCaches(context.Background(), prefix)
	if err != nil {
		return nil, err
	}

	products := make([]api.Product, 0, len(cached))
	for _, cache := range cached {
		product, err := convertToProduct(cache.Value)
		if err != nil {
			continue
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *productImpl) Update(cli client.Interface, name string, input *api.Product) (*api.Product, error) {
	if name != input.Name {
		return nil, errors.ErrorInvalidField.Error("name")
	}

	cache := new(v1alpha1.Cache)
	cache.Key = prefix + input.Name
	cache.Value = *input

	newCache, err := cli.V1Alpha1().UpdateCache(context.Background(), prefix+name, cache)
	if err != nil {
		return nil, err
	}

	product, err := convertToProduct(newCache.Value)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productImpl) Delete(cli client.Interface, name string) error {
	return cli.V1Alpha1().DeleteCache(context.Background(), prefix+name)
}

func (p *productImpl) Get(cli client.Interface, name string) (*api.Product, error) {
	cached, err := cli.V1Alpha1().GetCache(context.Background(), prefix+name)
	if err != nil {
		return nil, err
	}

	product, err := convertToProduct(cached.Value)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func convertToProduct(value interface{}) (api.Product, error) {
	buf, err := json.Marshal(value)
	if err != nil {
		return api.Product{}, err
	}

	var product api.Product
	err = json.Unmarshal(buf, &product)
	if err != nil {
		return api.Product{}, err
	}

	return product, nil
}
