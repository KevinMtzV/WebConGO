package controllers

import (
	"net/http"
	"proyecto-go/database"
	"proyecto-go/models"

	"github.com/gin-gonic/gin"
)

// Obtener todos los productos
func ObtenerProductos(c *gin.Context) {
	var productos []models.Producto
	database.DB.Find(&productos)
	c.JSON(http.StatusOK, productos)
}

// Obtener producto por ID
func ObtenerProducto(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto
	if err := database.DB.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, producto)
}

// Crear un nuevo producto
func CrearProducto(c *gin.Context) {
	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&producto)
	c.JSON(http.StatusCreated, producto)
}

// Actualizar un producto
func ActualizarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto
	if err := database.DB.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&producto)
	c.JSON(http.StatusOK, producto)
}

// Eliminar un producto
func EliminarProducto(c *gin.Context) {
	id := c.Param("id")
	var producto models.Producto
	if err := database.DB.First(&producto, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	database.DB.Delete(&producto)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Producto eliminado"})
}
