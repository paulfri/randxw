package main

import (
	"math/rand"
	"strings"
	"time"
)

// https://www.nytimes.com/crosswords/game/daily/1993/11/21
const urlFormat = "https://www.nytimes.com/crosswords/game/daily/%s"
const dateSegmentFormat = "2006/01/02"
const firstAvailableDateRFC3339 = "1993-11-21T00:00:00-07:00"

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
	firstAvailable, err := time.Parse(time.RFC3339, firstAvailableDateRFC3339)

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
