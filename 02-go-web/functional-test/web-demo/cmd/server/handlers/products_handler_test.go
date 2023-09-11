package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"web-demo/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// cases to test:
// - happy path								-> 200
// - handler: invalid product identifier	-> 400
// - service: product not found				-> 404
// - service: internal server error			-> 500

// Tests for ProductsHandler.GetByID
func CreateTestServerForProductsHandler(hd *ProductsHandler) *gin.Engine {
	r := gin.New()
	r.GET("/products/:id", hd.GetByID)
	return r
}

func TestFunctional_ProductsHandler_GetByID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// arrange
		// -> expectations
		expectedStatusCode := 200
		expectedResponseBody := `{"id":1,"name":"Galletitas Boreo","price":100}`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// -> setup
		rp := products.SliceBasedRepository{
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
		sv := products.DefaultService{Repository: &rp}
		h := ProductsHandler{Service: sv}
		server := CreateTestServerForProductsHandler(&h)

		// -> input
		request := httptest.NewRequest("GET", "/products/1", nil)
		response := httptest.NewRecorder()

		// act
		server.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectedResponseBody, response.Body.String())
		assert.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("handler: invalid product identifier", func(t *testing.T) {
		// arrange
		// -> expectations
		expectedStatusCode := 400
		expectedResponseBody := `{"error":"invalid product identifier"}`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// -> setup
		rp := products.SliceBasedRepository{}
		sv := products.DefaultService{Repository: &rp}
		h := ProductsHandler{Service: sv}
		server := CreateTestServerForProductsHandler(&h)

		// -> input
		request := httptest.NewRequest("GET", "/products/invalid", nil)
		response := httptest.NewRecorder()

		// act
		server.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectedResponseBody, response.Body.String())
		assert.Equal(t, expectedHeaders, response.Header())
	})

	t.Run("service: product not found", func(t *testing.T) {
		// arrange
		// -> expectations
		expectedStatusCode := 404
		expectedResponseBody := `{"error":"product not found"}`
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		// -> setup
		rp := products.SliceBasedRepository{}
		sv := products.DefaultService{Repository: &rp}
		h := ProductsHandler{Service: sv}
		server := CreateTestServerForProductsHandler(&h)

		// -> input
		request := httptest.NewRequest("GET", "/products/1", nil)
		response := httptest.NewRecorder()

		// act
		server.ServeHTTP(response, request)

		// assert
		assert.Equal(t, expectedStatusCode, response.Code)
		assert.JSONEq(t, expectedResponseBody, response.Body.String())
		assert.Equal(t, expectedHeaders, response.Header())
	})
}