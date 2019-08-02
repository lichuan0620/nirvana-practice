package handlers

import (
	"context"

	"github.com/lichuan0620/nirvana-practice/pkg/apis/middlewares"
	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

func CreateProduct(ctx context.Context, input *api.Product) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	productService, exist := middlewares.GetProductService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return productService.Create(client, input)
}

func ListProducts(ctx context.Context) ([]api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	productService, exist := middlewares.GetProductService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return productService.List(client)
}

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	productService, exist := middlewares.GetProductService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return productService.Get(client, name)
}

func UpdateProduct(ctx context.Context, name string, new *api.Product) (*api.Product, error) {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	productService, exist := middlewares.GetProductService(ctx)
	if !exist {
		return nil, errors.ErrorInternal.Error("")
	}

	return productService.Update(client, name, new)
}

func DeleteProduct(ctx context.Context, name string) error {
	client, exist := middlewares.GetClientInterface(ctx)
	if !exist {
		return errors.ErrorInternal.Error("")
	}

	productService, exist := middlewares.GetProductService(ctx)
	if !exist {
		return errors.ErrorInternal.Error("")
	}

	return productService.Delete(client, name)
}
