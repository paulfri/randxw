package main

import (
	"math/rand"
	"strings"
	"time"
)

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
