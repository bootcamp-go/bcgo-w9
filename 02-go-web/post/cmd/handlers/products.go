package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewControllerProducts(db []*Product, lastId int) *ControllerProducts {
	return &ControllerProducts{
		db: db,
		lastId: lastId,
	}
}
	

// Product is an struct that represents a product
type Product struct {
	Id			int
	Name		string
	Type		string
	Price		float64
	Quantity	int
}

func Validator(pr *Product) (err error) {
	// required
	if pr.Name == "" {
		err = errors.New("name is required")
		return
	}
	if pr.Type == "" {
		err = errors.New("type is required")
		return
	}

	return
}

// ControllerProducts is an struct that represents a controller for products
// exposing methods to handle products
type ControllerProducts struct {
	db 		[]*Product
	lastId	int
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
		pr := &Product{
			Name: reqBody.Name,
			Type: reqBody.Type,
			Price: reqBody.Price,
			Quantity: reqBody.Quantity,
		}
		pr.Id = c.lastId + 1
		// -> validation
		if err := Validator(pr); err != nil {
			code := http.StatusConflict
			body := ResponseBody{Message: "invalid product"}

			ctx.JSON(code, body)
			return
		}
		// -> save in storage
		c.db = append(c.db, pr)

		c.lastId++

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