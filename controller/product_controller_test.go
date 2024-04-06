package controller

import (
	"api-golang/dto"
	"api-golang/models"
	"api-golang/services/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {

	// Prepare de ProductService Interface
	mockService := new(mocks.ProductService)
	// Prepare de The Product Instance with data
	mockProduct := models.Product{Nombre: "Test Product", Precio: 10.99}
	// Prepare de The DTO with data to validate
	mockProductDTO := dto.ProductCreateDTO{Nombre: "Test Product", Precio: 10.99}

	// Mocking the service
	mockService.On("CreateProduct", mock.AnythingOfType("models.Product")).Return(mockProduct, nil)

	// Prepare the Controller and pass the Mock Service into the controller
	controller := NewProductController(mockService)

	// Http Request
	productJSON, _ := json.Marshal(mockProductDTO)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(productJSON))
	rr := httptest.NewRecorder()

	// After the http request call we use the controller to use the CreateProduct function
	handler := http.HandlerFunc(controller.CreateProduct)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the mock data vs response
	var returnedProduct models.Product
	json.Unmarshal(rr.Body.Bytes(), &returnedProduct)

	assert.Equal(t, mockProduct.Nombre, returnedProduct.Nombre)
	assert.Equal(t, mockProduct.Precio, returnedProduct.Precio)

	mockService.AssertExpectations(t)
}
