package controller

import (
	"api-golang/dto"
	"api-golang/models"
	"api-golang/services"
	"encoding/json"
	"net/http"
)

type ProductController struct {
	Service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	productos, err := c.Service.GetAllProductos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(productos)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.ProductCreateDTO
	if err := json.NewDecoder(r.Body).Decode(&productDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// DTO Validation
	if validationErrors := productDTO.Validate(); validationErrors != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	product := models.Product{Nombre: productDTO.Nombre, Precio: productDTO.Precio}
	savedProduct, err := c.Service.CreateProduct(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(savedProduct)
}
