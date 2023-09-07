package main

import (
	"app/cmd/handlers"
	"app/internal/products/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// dependencies
	stMe := storage.NewStorageProductsInMemory(make([]*storage.Product, 0), 0)
	stVl := storage.NewStorageProductsValidation(stMe)
	ct := handlers.NewControllerProducts(stVl)

	// server
	rt := gin.Default()
	// -> middlewares
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())
	// -> routes
	rt.POST("/products", ct.Save())

	// run
	if err := rt.Run(":8080"); err != nil {
		panic(err)
	}
}