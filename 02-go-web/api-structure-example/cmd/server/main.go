package main

import (
	"app/cmd/server/handlers"
	"app/internal/domain"
	"app/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// dependencies
	rp := products.NewRepositoryProductInMemory(make([]*domain.Product, 0))
	sv := products.NewServiceProducts(rp)
	hd := handlers.NewHandlerProducts(sv)

	// server
	rt := gin.Default()
	// -> endpoints
	rt.GET("/products/:id", hd.GetByID())

	// run
	if err := rt.Run(":8080"); err != nil {
		panic(err)
	}
}