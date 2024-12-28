package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Item 1"},
	{ID: "2", Name: "Item 2"},
}

func main() {
	// Crear una nueva instancia del router
	router := gin.Default()

	router.GET("/items", func(c *gin.Context) {
		c.JSON(http.StatusOK, items)
	})

	// Endpoint POST para crear un nuevo item
	router.POST("/items", func(c *gin.Context) {
		var newItem Item

		// Enlazar el JSON recibido al objeto newItem
		if err := c.ShouldBindJSON(&newItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Agregar el nuevo item a la lista
		items = append(items, newItem)
		c.JSON(http.StatusCreated, newItem)
	})

	// Inicia el servidor en el puerto 8080
	router.Run(":8080")
}
