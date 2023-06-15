package http

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CreateProductRequest struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
}

func (r *CreateProductRequest) Decode(req *http.Request) error {
	return json.NewDecoder(req.Body).Decode(r)
}

func (r *CreateProductRequest) Validate() error {
	if r.Category == "" {
		return errors.New("category is required")
	}
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.SKU == "" {
		return errors.New("sku is required")
	}
	return nil
}

type CreateProductResponse struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
}
