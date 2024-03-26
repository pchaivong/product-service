package main

import (
	"github.com/pchaivong/product-service/internal/adapters/handlers/rest"
	"github.com/pchaivong/product-service/internal/adapters/repository"
	"github.com/pchaivong/product-service/internal/core/service"
)

func main() {

	r := repository.NewCouchDBRepository()
	s := service.NewProductService(r)
	h := rest.NewGinHTTPHandler(s, &rest.GinHandlerOpt{Addr: ":8080"})

	err := h.Start()
	if err != nil {
		panic(err)
	}
}
