package main

import (
	"proyecto-go/database"
	"proyecto-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	database.ConectarBD()

	// Inicializar router
	r := gin.Default()

	// Servir archivos estáticos y templates
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Página principal
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Rutas de productos
	routes.ProductoRoutes(r)

	// Iniciar servidor
	r.Run(":8080")
}
