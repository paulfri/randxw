package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*.html")

	r.GET("/", todayCrosswordRoute)
	r.GET("/random", randomCrosswordRoute)
	r.GET("/:dow", dowCrosswordRoute)
	r.GET("/clue", searchRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
