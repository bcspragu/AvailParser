package avail

import (
	"strconv"
	"strings"
	"time"
)

// Makes a legible message from a query
func HumanReadableMessage(result QueryResult) string {
	currentTime := time.Now()
	mSegs := make([]string, 0, len(result))

	for stop, depsByRoute := range result {
		mSegs = append(mSegs, "At stop "+stop.Name+":")
		for route, deps := range depsByRoute {
			for _, departure := range deps {
				mSegs = append(mSegs, "*"+route.ShortName+" at "+departure.Trip.InternetServiceDesc+" in "+durationAsAMessage(departure.EDT.Sub(currentTime)))
			}
			mSegs = append(mSegs, "----------")
		}
	}

	return strings.Join(mSegs, "\n")
}

// Turn a duration into a semi-readable string
func durationAsAMessage(duration time.Duration) string {
	mSegs := make([]string, 0, 3)

	seconds := int(duration.Seconds()) % 60
	minutes := int(duration.Minutes()) % 60
	hours := int(duration.Hours())

	if hours > 0 {
		switch hours {
		case 1:
			mSegs = append(mSegs, "an hour")
		default:
			mSegs = append(mSegs, strconv.Itoa(int(hours))+" hours")
		}
	}

	if minutes > 0 {
		switch minutes {
		case 1:
			mSegs = append(mSegs, "a minute")
		default:
			mSegs = append(mSegs, strconv.Itoa(int(minutes))+" minutes")
		}
	}

	if seconds > 0 {
		mSegs = append(mSegs, strconv.Itoa(int(seconds))+" second(s)")
	}
	return strings.Join(mSegs, ", ")
}
