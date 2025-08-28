package main

import (
	"fmt"
	"temp/gin-service/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	// router.GET("/albums", func(c *gin.Context) {
	// 	albums := handler.GetAlbums()
	// 	c.IndentedJSON(http.StatusOK, albums)
	// })
	router.GET("/getAlbums", handler.GetAlbums)
	router.POST("/createalbum", handler.AddAlbum)
	router.Run(":4040")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// var a bool
	// fmt.Println("Enter a number:")
	// fmt.Scanln(&a)
	// if a == false {
	// 	panic("You entered false, which is not allowed!")
	// } else {
	// 	fmt.Println("You entered true, proceeding with the program.")
	// }

}
