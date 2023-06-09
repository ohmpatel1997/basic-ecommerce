package http

import (
	"encoding/json"
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

type CreateProductResponse struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
}
