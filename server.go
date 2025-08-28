package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"temp/gin-service/com/handler"
	"time"

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
	router.Run("localhost:8080")

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

	var b []int = []int{1, 2, 3, 4, 5}
	b = append(b, 1)
	fmt.Println("Array b:", b)
	fmt.Println("\"Program completed successfully.\"")

	var c string = `
	Hello, World!
	Welcome to the Go programming language.
	Enjoy coding!`
	fmt.Println(c)

	strDecimal := "456"
	integerDecimal, err := strconv.ParseInt(strDecimal, 10, 32)
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}
	fmt.Printf("Converted decimal to int: %d, type: %T\n", integerDecimal, integerDecimal)
	type Person struct {
		Name string `json:"name,omitempty"`
		age  int
	}
	p1 := &Person{Name: "dss", age: 25}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	} else {
		fmt.Println("JSON data:", string(jsonData))
	}

	ptrP := new(Person)
	err = json.Unmarshal(jsonData, ptrP)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		//fmt.Println("JSON data:", string(jsonData))
	}

	fmt.Println("Pointer to Person struct:", ptrP)
	fmt.Printf("Pointer to Person struct 1: %+v", *p1)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(1500 * time.Millisecond): // Timeout case
		fmt.Println("Timeout: No channel ready within 1.5 seconds")
	}

}
