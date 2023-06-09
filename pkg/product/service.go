package product

import (
	"context"

	"github.com/ohmpatel1997/rundoo-task/pkg/storage"
	"github.com/ohmpatel1997/rundoo-task/pkg/storage/types"
)

type ServiceI interface {
	CreateProduct(ctx context.Context, name, SKU, category string) (*types.Product, error)
	SearchProduct(ctx context.Context, query string, limit int64) ([]*types.Product, error)
}

type Service struct {
	storage storage.StorageI
}

func NewService(storage storage.StorageI) ServiceI {
	return &Service{storage: storage}
}

func (s *Service) CreateProduct(ctx context.Context, name, SKU, category string) (*types.Product, error) {
	product := types.NewProduct(
		types.WithName(name),
		types.WithSKU(SKU),
		types.WithCategory(category),
	)
	return s.storage.CreateProduct(ctx, product)
}

func (s *Service) SearchProduct(ctx context.Context, query string, limit int64) ([]*types.Product, error) {
	return s.storage.SearchProduct(ctx, query, limit)
}
