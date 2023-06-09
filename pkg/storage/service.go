package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ohmpatel1997/rundoo-task/pkg/storage/types"
)

type StorageI interface {
	Close()
	CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error)
	SearchProduct(ctx context.Context, query string, limit int64) ([]*types.Product, error)
}

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) StorageI {
	return &Service{db: db}
}

func (s *Service) Close() {
	s.db.Close()
}

func (s *Service) CreateProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	prod := new(types.Product)
	err := s.db.QueryRow(ctx, "INSERT INTO products (category, name, sku) VALUES ($1, $2, $3)  returning id, category, name, sku", product.Category, product.Name, product.SKU).
		Scan(&product.ID, &product.Category, &product.Name, &product.SKU)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (s *Service) SearchProduct(ctx context.Context, query string, limit int64) ([]*types.Product, error) {
	rows, err := s.db.Query(
		ctx,
		"SELECT * FROM products WHERE to_tsquery($1) @@ to_tsvector(courses.category || courses.name || courses.sku) LIMIT $2",
		query,
		limit,
	)

	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		var product *types.Product
		err = rows.Scan(&product.ID, &product.Category, &product.Name, &product.SKU)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return products, nil
}
