package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Knife", Artist: "ABSL", Price: 21.00},
	{ID: "2", Title: "GRND", Artist: "Hioll", Price: 21.00},
	{ID: "3", Title: "Abbdon", Artist: "RubberMind", Price: 21.00},
}

func getAlbums(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(g *gin.Context) {
	id := g.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, v := range albums {
		if v.ID == id {
			g.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func postAlbum(p *gin.Context) {
	var newAlbum Album

	if err := p.BindJSON(&newAlbum); err != nil {
		return
	}

	/* Add the new album to the slice */
	albums = append(albums, newAlbum)
	p.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8000")
}
