package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	time := timeInNewYork()
	results := scrapeAnswers(query, time)

	c.HTML(http.StatusOK, "results.html", gin.H{
		"query":   query,
		"date":    dateString(time),
		"results": results,
	})
}
