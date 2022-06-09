package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

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
