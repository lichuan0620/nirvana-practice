package handlers

import (
	"context"

	api "github.com/lichuan0620/nirvana-practice/pkg/apis/v1alpha1"
	"github.com/lichuan0620/nirvana-practice/pkg/errors"
)

func GetProduct(ctx context.Context, name string) (*api.Product, error) {
	return nil, errors.ErrorNotImplemented.Error()
}
