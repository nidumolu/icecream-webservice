package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type icecream struct {
	ID     string  `json:"id"`
	Flavor string  `json:"flavor"`
	Price  float64 `json:"price"`
}

var icecreams = []icecream{
	{ID: "1", Flavor: "Vanilla", Price: 1.99},
	{ID: "2", Flavor: "Chocolate", Price: 2.99},
	{ID: "3", Flavor: "Strawberry", Price: 2.99},
}

func getIcecreams(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, icecreams)
}

func postIcecreams(c *gin.Context) {
	var newIcecream icecream
	// Call BindJSON to bind the received JSON to
	// newIcecream.
	if err := c.BindJSON(&newIcecream); err != nil {
		return
	}

	// Add the new icecream to the slice.
	icecreams = append(icecreams, newIcecream)
	c.IndentedJSON(http.StatusCreated, newIcecream)
}

// getIcecreamByID locates the icecream whose ID value matches the id
// parameter sent by the client, then returns that icecream as a response.
func getIcecreamByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of icecreams, looking for
	// an icecream whose ID value matches the parameter.
	for _, a := range icecreams {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "That icecream was not found :-( "})
}

func main() {
	router := gin.Default()
	router.GET("/icecreams", getIcecreams)
	router.POST("/icecreams", postIcecreams)
	router.GET("/icecreams/:id", getIcecreamByID)
	router.Run("localhost:8080")
}
