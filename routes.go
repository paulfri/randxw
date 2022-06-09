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
