package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// https://www.nytimes.com/crosswords/game/daily/1993/11/21
const urlFormat = "https://www.nytimes.com/crosswords/game/daily/%s"
const dateSegmentFormat = "2006/01/02"

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	r.GET("/", todayCrosswordRoute)
	r.GET("/random", randomCrosswordRoute)
	r.GET("/:dow", dowCrosswordRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
