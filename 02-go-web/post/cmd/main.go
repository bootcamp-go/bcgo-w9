package main

import (
	"app/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// env
	// ...

	// dependencies
	db := make([]*handlers.Product, 0)
	ct := handlers.NewControllerProducts(db, 0)

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