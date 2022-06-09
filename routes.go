package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func todayCrosswordRoute(c *gin.Context) {
	c.Redirect(http.StatusFound, todayCrosswordURL())
}

func randomCrosswordRoute(c *gin.Context) {
	c.Redirect(http.StatusFound, randomCrosswordURL())
}

func dowCrosswordRoute(c *gin.Context) {
	dow := c.Param("dow")

	if isValidDow(dow) {
		c.Redirect(http.StatusFound, randomCrosswordURLByDayOfWeek(dow))
	} else {
		c.Status(http.StatusNotFound)
	}
}

func searchRoute(c *gin.Context) {
	query := c.Query("q")

	results := scrapeAnswers(query)

	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"results": results,
	})
}

const searchURLFormat = "https://www.xwordinfo.com/Crossword?date=%s"

func scrapeAnswers(query string) []string {
	var matches []string

	lowerQuery := strings.ToLower(query)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#CPHContent_ClueBox .numclue div", func(e *colly.HTMLElement) {
		fmt.Println("Found element:", e.Text)

		content := e.Text
		lowerContent := strings.ToLower(content)

		if strings.Contains(lowerContent, lowerQuery) {
			matches = append(matches, content)
		}
	})

	today := time.Now()
	answerURL := fmt.Sprintf(searchURLFormat, dateString(today))
	c.Visit(answerURL)

	return matches
}
