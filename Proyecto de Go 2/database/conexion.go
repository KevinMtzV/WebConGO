package database

import (
	"fmt"
	"log"
	"proyecto-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBD() {
	dsn := "root:RootKevin@tcp(localhost:3306)/tienda?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	// Migración de modelos
	DB.AutoMigrate(&models.Producto{})

	fmt.Println("✅ Conectado a la base de datos")
}
