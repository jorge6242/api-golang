package repositories

import (
	"api-golang/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	Save(producto models.Product) (models.Product, error)
}

// Struct a el ORM Gorm a fin de usar las bondades del orm
type productoRepository struct {
	db *gorm.DB
}

func NewProductoRepository(db *gorm.DB) ProductRepository {
	return &productoRepository{db: db}
}

func (r *productoRepository) FindAll() ([]models.Product, error) {
	var productos []models.Product
	result := r.db.Find(&productos)
	return productos, result.Error
}

func (r *productoRepository) Save(producto models.Product) (models.Product, error) {
	result := r.db.Create(&producto)
	return producto, result.Error
}
