package handler

import (
	"net/http"
	"temp/gin-service/model"

	"github.com/gin-gonic/gin"
)

var albums = []model.Album{
	{
		ID:     1,
		Title:  "Album Title",
		Author: "Artist Name",
		Price:  9.99,
	},
	{
		ID:     2,
		Title:  "Second Album",
		Author: "Another Artist",
		Price:  12.99,
	},
}

func GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbum(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
	} else {
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, albums)
	}

}
