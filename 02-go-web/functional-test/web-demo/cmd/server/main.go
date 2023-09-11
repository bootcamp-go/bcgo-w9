package main

import (
	"web-demo/cmd/server/handlers"
	"web-demo/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	productsRepository := &products.SliceBasedRepository{
		Data: []*products.Product{
			{
				ID:    1,
				Name:  "Galletitas Boreo",
				Price: 100,
			},
			{
				ID:    2,
				Name:  "Moca Cola",
				Price: 200,
			},
		},
	}
	productsService := products.DefaultService{
		Repository: productsRepository,
	}
	productsHandler := handlers.ProductsHandler{
		Service: productsService,
	}

	// Create a new Gin router
	router := gin.Default()

	// Create the routes.
	productsRoutes := router.Group("/products")
	{
		productsRoutes.GET("/:id", productsHandler.GetByID)
	}
}
