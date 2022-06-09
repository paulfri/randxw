package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

const searchURLFormat = "https://www.xwordinfo.com/Crossword?date=%s"

func scrapeAnswers(query string, date time.Time) []fuzzy.Rank {
	var answers []string

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#CPHContent_ClueBox .numclue div", func(e *colly.HTMLElement) {
		fmt.Println("Found element:", e.Text)

		answers = append(answers, e.Text)
	})

	answerURL := fmt.Sprintf(searchURLFormat, dateString(date))
	c.Visit(answerURL)

	matches := fuzzy.RankFindNormalizedFold(query, answers)
	sort.Sort(matches)

	return matches
}
