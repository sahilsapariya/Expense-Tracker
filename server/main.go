package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: 1, Name: "Item One"},
    {ID: 2, Name: "Item Two"},
}

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API!"})
    })

    r.GET("/items", func(c *gin.Context) {
        c.JSON(http.StatusOK, items)
    })

    r.POST("/items", func(c *gin.Context) {
        var newItem Item
        if err := c.ShouldBindJSON(&newItem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newItem.ID = len(items) + 1
        items = append(items, newItem)
        c.JSON(http.StatusCreated, newItem)
    })

    r.PUT("/items/:id", func(c *gin.Context) {
        var updated Item
        if err := c.ShouldBindJSON(&updated); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        for i, item := range items {
            if item.ID == updated.ID {
                items[i].Name = updated.Name
                c.JSON(http.StatusOK, items[i])
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
    })

    r.DELETE("/items/:id", func(c *gin.Context) {
        id := c.Param("id")
        for i, item := range items {
            if id == string(rune(item.ID)) {
                items = append(items[:i], items[i+1:]...)
                c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
    })

    r.Run(":8080")
}
