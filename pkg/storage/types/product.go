package types

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id"`
	Category string    `json:"category"`
	Name     string    `json:"name"`
	SKU      string    `json:"sku"`
}

type CreateProductOpt func(*Product)

func NewProduct(opts ...CreateProductOpt) *Product {
	p := &Product{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithCategory(category string) CreateProductOpt {
	return func(p *Product) {
		p.Category = category
	}
}

func WithName(name string) CreateProductOpt {
	return func(p *Product) {
		p.Name = name
	}
}

func WithSKU(sku string) CreateProductOpt {
	return func(p *Product) {
		p.SKU = sku
	}
}
