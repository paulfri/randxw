package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// https://www.nytimes.com/crosswords/game/daily/1993/11/21
const format = "https://www.nytimes.com/crosswords/game/daily/%s"

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, randomCrosswordURL())
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func randomCrosswordURL() string {
	date := randomDate()
	dateSegment := dateString(date)

	return fmt.Sprintf(format, dateSegment)
}

func randomDate() time.Time {
	nowUnix := time.Now().Unix()
	firstAvailable, err := time.Parse(time.RFC3339, "1993-11-21T00:00:00-07:00")

	if err != nil {
		panic(err)
	}

	firstAvailableUnix := firstAvailable.Unix()
	randomTimestamp := rand.Int63n(nowUnix-firstAvailableUnix) + firstAvailableUnix
	randomUnix := time.Unix(randomTimestamp, 0)

	return randomUnix
}

func dateString(unix time.Time) string {
	return unix.Format("2006/01/02")
}
