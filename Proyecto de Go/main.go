package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Conexión a la base de datos
	dsn := "root:RootKevin@tcp(localhost:3306)/tienda?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}

	// Migración automática del modelo
	db.AutoMigrate(&Producto{})

	// Inicializar router
	r := gin.Default()

	// Rutas de la API
	r.GET("/productos", obtenerProductos)
	r.GET("/productos/:id", obtenerProducto)
	r.POST("/productos", crearProducto)
	r.PUT("/productos/:id", actualizarProducto)
	r.DELETE("/productos/:id", eliminarProducto)

	// Servir archivos estáticos (Frontend)
	r.Static("/static", "./static")

	// Página principal
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Iniciar servidor
	r.Run(":8080")
}

// ==================== CONTROLADORES CRUD ====================

// Obtener todos los productos
func obtenerProductos(c *gin.Context) {
	var productos []Producto
	db.Find(&productos)
	c.JSON(http.StatusOK, productos)
}

// Obtener producto por ID
func obtenerProducto(c *gin.Context) {
	id := c.Param("id")
	var producto Producto
	if err := db.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

// Crear un nuevo producto
func crearProducto(c *gin.Context) {
	var producto Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&producto)
	c.JSON(http.StatusCreated, producto)
}

// Actualizar un producto
func actualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto Producto
	if err := db.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&producto)
	c.JSON(http.StatusOK, producto)
}

// Eliminar un producto
func eliminarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto Producto
	if err := db.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	db.Delete(&producto)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado"})
}

//para iniciar el codigo:
//go run main.go productos.go
