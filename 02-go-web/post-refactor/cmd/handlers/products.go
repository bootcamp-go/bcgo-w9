package handlers

import (
	"app/internal/products/storage"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewControllerProducts is a function that returns a new controller of products
func NewControllerProducts(st storage.StorageProducts) *ControllerProducts {
	return &ControllerProducts{
		st: st,
	}
}

// ControllerProducts is an struct that returns handlers of products
type ControllerProducts struct {
	st storage.StorageProducts
}

// Save is a method that handles the save of a product
type RequestBody struct {
	Name		string	`json:"name" binding:"required"`
	Type		string	`json:"type" binding:"required"`
	Price		float64	`json:"price" binding:"required"`
	Quantity	int		`json:"quantity"`
}
type Data struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Type		string	`json:"type"`
	Price		float64	`json:"price"`
	Quantity	int		`json:"quantity"`
}
type ResponseBody struct {
	Message		string	`json:"message"`
	Data		*Data	`json:"data"`
}
func (c *ControllerProducts) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// -> header
		token := ctx.GetHeader("Authorization")
		if token != "123" {
			code := http.StatusUnauthorized
			body := ResponseBody{Message: "Invalid token"}

			ctx.JSON(code, body)
			return
		}

		// -> body
		var reqBody RequestBody
		err := ctx.ShouldBindJSON(&reqBody)
		if err != nil {
			code := http.StatusBadRequest
			body := ResponseBody{Message: "Invalid request body"}

			ctx.JSON(code, body)
			return
		}
		
		// process
		// -> deserialization
		pr := &storage.Product{
			Name: reqBody.Name,
			Type: reqBody.Type,
			Price: reqBody.Price,
			Quantity: reqBody.Quantity,
		}
		// -> save
		err = c.st.Save(pr)
		if err != nil {
			var code int; var body ResponseBody
			switch {
				case errors.Is(err, storage.ErrStorageProductsInvalid):
					code = http.StatusConflict
					body = ResponseBody{Message: "Invalid product"}
				default:
					code = http.StatusInternalServerError
					body = ResponseBody{Message: "Internal error"}
			}

			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusCreated		
		body := ResponseBody{
			Message: "Product created",
			Data: &Data{				// serialization
				Id: pr.Id,
				Name: pr.Name,
				Type: pr.Type,
				Price: pr.Price,
				Quantity: pr.Quantity,
			},
		}

		ctx.JSON(code, body)
	}
}