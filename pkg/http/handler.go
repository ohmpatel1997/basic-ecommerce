package http

import (
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"github.com/ohmpatel1997/rundoo-task/pkg/product"
)

type Handler struct {
	productSvc product.ServiceI
}

func NewHandler(productSvc product.ServiceI) *Handler {
	return &Handler{productSvc: productSvc}
}

func (h *Handler) RegisterRoutes(r chi.Router) http.Handler {
	r.Use(middleware.Heartbeat("/ping"))

	cors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		Debug:            false,
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(panicRecover)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/products", func(r chi.Router) {
			r.Post("/", h.creatProduct)
			r.Get("/search", h.searchProducts)
		})
	})

	return r
}

func (h *Handler) creatProduct(w http.ResponseWriter, r *http.Request) {
	req := new(CreateProductRequest)
	if err := req.Decode(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	invite, err := h.productSvc.CreateProduct(r.Context(), req.Name, req.SKU, req.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, invite)
}

func (h *Handler) searchProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, errors.New("invalid query param").Error(), http.StatusInternalServerError)
		return
	}

	limit := r.URL.Query().Get("limit")
	if len(limit) == 0 {
		limit = "10"
	}

	intLimit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 5 <= limit <= 50
	intLimit = int64(math.Max(5, float64(intLimit)))
	intLimit = int64(math.Min(50, float64(intLimit)))

	products, err := h.productSvc.SearchProduct(r.Context(), query, intLimit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, products)
}
