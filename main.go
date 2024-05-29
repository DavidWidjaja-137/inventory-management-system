package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type InventoryItem struct {
	ItemName        string    `json:"name"`
	ItemDescription string    `json:"description"`
	ObtainedFrom    string    `json:"obtained_from"`
	InPossession    bool      `json:"in_possession"`
	ItemType        ItemType  `json:"item_type"`
	DateAcquired    time.Time `json:"date_acquired"`
	DateReleased    time.Time `json:"date_released"`
}

var items = []InventoryItem{}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, make([]InventoryItem, 0))
}

func postItems(c *gin.Context) {
	var newItem InventoryItem

	if err := c.BindJSON(&newItem); err != nil {
		return
	}

	items = append(items, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)

}

func main() {

	fmt.Println("Starting Server...")

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	router.Use(cors.New(config))

	router.GET("/items/", getItems)
	router.POST("/items/", postItems)

	router.Run("localhost:8080")
}
