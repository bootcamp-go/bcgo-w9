package handlers

import (
	"net/http"
	"strconv"
	"web-demo/internal/products"

	"github.com/gin-gonic/gin"
)

type GetByIDResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductsHandler struct {
	Service products.Service
}

func (handler ProductsHandler) GetByID(ctx *gin.Context) {
	// Get the ID from the URL.
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product identifier",
		})
		return
	}

	// Get the product from the service.
	product, err := handler.Service.GetByID(id)
	if err != nil {
		if err == products.ErrProductNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "product not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	// Return the product.
	ctx.JSON(http.StatusOK, GetByIDResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	})
}
