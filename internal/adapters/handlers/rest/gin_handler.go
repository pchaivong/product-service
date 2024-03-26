package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pchaivong/product-service/internal/core/domain"
	"github.com/pchaivong/product-service/internal/core/ports"
)

type GinHandlerOpt struct {
	Addr string
}

type ginhandler struct {
	addr string
	s    ports.ProductService
	e    *gin.Engine
}

func NewGinHTTPHandler(s ports.ProductService, o *GinHandlerOpt) HTTPHandler {

	handler := &ginhandler{
		addr: o.Addr,
		s:    s,
		e:    gin.Default(),
	}

	handler.setup()
	return handler
}

func (h *ginhandler) Start() error {
	return h.e.Run(h.addr)
}

func (h *ginhandler) setup() {
	v1 := h.e.Group("v1")

	v1.POST("products", h.handleCreateProduct)
	v1.GET("products", h.handleGetAllProducts)
	v1.GET("products/:productId", h.handleGetProductById)
	v1.PUT("products/:productId", h.handleUpdateProduct)

}

// POST /v1/products
func (h *ginhandler) handleCreateProduct(c *gin.Context) {
	var req CreateOrUpdateProductRequest
	err := c.BindJSON(&req)
	if err != nil {
		// Log something for debuging
		c.JSON(http.StatusBadRequest, ErrBadRequest)
		return
	}

	p, err := h.s.CreateProduct(req.Name, req.SKU, req.Description, req.Price, req.Available)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	res := ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		SKU:         p.SKU,
		Description: p.Description,
		Price:       p.Price,
		Available:   p.Available,
	}

	c.JSON(http.StatusCreated, res)
}

// GET /v1/products
func (h *ginhandler) handleGetAllProducts(c *gin.Context) {
	resp := make([]ProductResponse, 0)
	products := h.s.ListProduct()
	for _, p := range products {
		dto := ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			SKU:         p.SKU,
			Price:       p.Price,
			Available:   p.Available,
		}

		resp = append(resp, dto)
	}
	c.JSON(http.StatusOK, resp)
}

// GET /v1/products/{productId}
func (h *ginhandler) handleGetProductById(c *gin.Context) {
	id := c.Param("productId")
	p, err := h.s.GetProduct(id)
	if err == ports.ErrProductNotFound {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, p)
}

// PUT /v1/products/{productId}
func (h *ginhandler) handleUpdateProduct(c *gin.Context) {
	id := c.Param("productId")
	var req CreateOrUpdateProductRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrBadRequest)
		return
	}
	p := domain.Product{
		ID:          id,
		Name:        req.Name,
		SKU:         req.SKU,
		Description: req.Description,
		Price:       req.Price,
		Available:   req.Available,
	}

	err = h.s.UpdateProduct(&p)
	if err == ports.ErrProductNotFound {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "product not found",
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, p)
}
