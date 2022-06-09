package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// https://www.nytimes.com/crosswords/game/daily/1993/11/21
const urlFormat = "https://www.nytimes.com/crosswords/game/daily/%s"
const dateSegmentFormat = "2006/01/02"

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, todayCrosswordURL())
	})

	r.GET("/random", func(c *gin.Context) {
		c.Redirect(http.StatusFound, randomCrosswordURL())
	})

	r.GET("/:dow", func(c *gin.Context) {
		dow := c.Param("dow")

		if isValidDow(dow) {
			c.Redirect(http.StatusFound, randomCrosswordURLByDayOfWeek(dow))
		} else {
			c.Status(http.StatusNotFound)
		}
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

	return fmt.Sprintf(urlFormat, dateSegment)
}

func randomCrosswordURLByDayOfWeek(dow string) string {
	var generatedDow string
	var date time.Time
	for ok := true; ok; ok = (strings.ToLower(generatedDow) != strings.ToLower(dow)) {
		date = randomDate()
		generatedDow = date.Format("Mon")
		log.Println(generatedDow)
	}

	dateSegment := dateString(date)

	return fmt.Sprintf(urlFormat, dateSegment)
}

func todayCrosswordURL() string {
	newYork, _ := time.LoadLocation("America/New_York")
	date := time.Now().In(newYork)
	dateSegment := dateString(date)

	return fmt.Sprintf(urlFormat, dateSegment)
}

func isValidDow(dow string) bool {
	switch strings.ToLower(dow) {
	case
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sat",
		"sun":
		return true
	}
	return false
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
	return unix.Format(dateSegmentFormat)
}
