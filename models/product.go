package models

type Product struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
