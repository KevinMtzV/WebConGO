package main

type Producto struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Nombre string  `gorm:"size:100;not null" json:"nombre"`
	Precio float64 `gorm:"not null" json:"precio"`
}
