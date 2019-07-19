package handlers

import (
	"context"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/middlewares"
	"github.com/lichuan0620/nirvana-practice/pkg/apis/service"
	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

func init() {
	ps = service.NewProductService()
}

var ps service.ProductService

func CreateProduct(ctx context.Context, input *api.Product) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return ps.Create(client, input)
}

func ListProducts(ctx context.Context) ([]api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return ps.List(client)
}

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return ps.Get(client, name)
}

func UpdateProduct(ctx context.Context, name string, new *api.Product) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return ps.Update(client, name, new)
}

func DeleteProduct(ctx context.Context, name string) error {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return errors.ErrorInternal.Error("")
	}

	return ps.Delete(client, name)
}
