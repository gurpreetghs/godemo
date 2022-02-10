package main

import (
	"log"

	"github.com/gin-gonic/gin"
)


func homepage(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
		"age":     26,
	})
}

func PostHomepage()

func main() {
	r := gin.Default()
	r.GET("/ping", homepage )
	if err := Run("8080"); err != nill{
		log.Fatal(err.Error())
	}	
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
