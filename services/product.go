package services

import (
	"api-golang/errors"
	"api-golang/models"
	"api-golang/repositories"
)

type ProductService interface {
	GetAllProductos() ([]models.Product, error)
	CreateProduct(producto models.Product) (models.Product, error)
}

type productoService struct {
	repository repositories.ProductRepository
}

func NewProductoService(repository repositories.ProductRepository) ProductService {
	return &productoService{repository: repository}
}

func (s *productoService) GetAllProductos() ([]models.Product, error) {
	res, err := s.repository.FindAll()
	if err != nil {
		return []models.Product{}, errors.New(errors.DatabaseError, "Error find all product: "+err.Error())
	}
	return res, nil
}

func (s *productoService) CreateProduct(producto models.Product) (models.Product, error) {
	// Intenta guardar el producto y maneja tanto el producto guardado como el error
	savedProduct, err := s.repository.Save(producto)
	if err != nil {
		// Aquí manejas el error, creando un nuevo AppError con tipo DatabaseError
		return models.Product{}, errors.New(errors.DatabaseError, "Error saving product: "+err.Error())
	}
	// Si no hay error, devuelve el producto guardado y nil para el error
	return savedProduct, nil
}
