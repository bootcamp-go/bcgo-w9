package handlers

import (
	"app/internal/products"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NewHandlerProducts creates a new handler for products
func NewHandlerProducts(sv *products.ServiceProducts) *HandlerProducts {
	return &HandlerProducts{sv: sv}
}

// HandlerProducts is a handler for products
type HandlerProducts struct {
	sv *products.ServiceProducts
}

// GetByID returns a product by id
type ResponseBodyGetByID struct {
	Message string `json:"message"`
	Data	*struct {
		ID int 		  `json:"id"`
		Name string	  `json:"name"`
		Price float64 `json:"price"`
	} `json:"data"`
}
func (hd *HandlerProducts) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			code := http.StatusBadRequest
			body := ResponseBodyGetByID{Message: "Invalid ID", Data: nil}

			c.JSON(code, body)
			return
		}

		// process
		pr, err := hd.sv.GetByID(id)
		if err != nil {
			var code int; var body ResponseBodyGetByID
			switch {
			case errors.Is(err, products.ErrRepositoryProductNotFound):
				code = http.StatusNotFound
				body = ResponseBodyGetByID{Message: "Product not found", Data: nil}
			default:
				code = http.StatusInternalServerError
				body = ResponseBodyGetByID{Message: "Internal server error", Data: nil}
			}

			c.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := ResponseBodyGetByID{
			Message: "Success",						// serialize from service model to response model
			Data: &struct {
				ID int 		  `json:"id"`
				Name string	  `json:"name"`
				Price float64 `json:"price"`
			}{ID: pr.Id, Name: pr.Name, Price: pr.Price},
		}

		c.JSON(code, body)
	}
}