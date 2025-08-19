package routes

import (
	"proyecto-go/controllers"

	"github.com/gin-gonic/gin"
)

func ProductoRoutes(r *gin.Engine) {
	productos := r.Group("/productos")
	{
		productos.GET("/", controllers.ObtenerProductos)
		productos.GET("/:id", controllers.ObtenerProducto)
		productos.POST("/", controllers.CrearProducto)
		productos.PUT("/:id", controllers.ActualizarProducto)
		productos.DELETE("/:id", controllers.EliminarProducto)
	}
}
