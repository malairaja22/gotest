package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	// var dir, file string

	// dir, file = path.Split("css/main.js")
	// fmt.Println("dir:", dir)
	// fmt.Println("file:", file)

	a1 := 5
	a2 := 2.5

	fmt.Println("res", a2*float64(a1))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	// myProgramName := os.Args[1]

	// // it will display
	// // the program name
	// fmt.Println(myProgramName)

}
