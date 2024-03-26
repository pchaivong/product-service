package rest

import "net/http"

type HTTPHandler interface {
	Start() error
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	SKU         string  `json:"sku"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	Available   bool    `json:"available"`
}

type CreateOrUpdateProductRequest struct {
	Name        string  `json:"name"`
	SKU         string  `json:"sku"`
	Description string  `json:"desc"`
	Price       float32 `json:"price"`
	Available   bool    `json:"available"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrBadRequest          = ErrorResponse{Code: http.StatusBadRequest, Message: "Bad Request"}
	ErrInternalServerError = ErrorResponse{Code: http.StatusInternalServerError, Message: "Internal Server Error"}
)
