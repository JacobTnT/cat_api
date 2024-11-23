package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cat struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

var cats = []Cat{
	{
		Name:  "michi",
		Breed: "gray",
	},
	{
		Name:  "tiger",
		Breed: "orange",
	},
}

// getCats return the list of cats in the server
func getCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

func getCat(c *gin.Context) {
	name := c.Param("name")

	for _, cat := range cats {
		if cat.Name == name {
			c.IndentedJSON(http.StatusOK, cat)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
}

func main() {
	router := gin.Default()
	router.GET("/cats", getCats)
	router.GET("/cats/:name", getCat)

	router.Run("localhost:8080")
}
