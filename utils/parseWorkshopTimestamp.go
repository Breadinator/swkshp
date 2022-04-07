package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// e.g. "25 Sep, 2021 @ 11:17am"
var timestampRegexp *regexp.Regexp = regexp.MustCompile(`(\d+) (\w+), (\d+) @ (\d+):(\d+)(\w+)`)

func ParseWorkshopTimestamp(url string) (time.Time, bool) {
	timestamp := "25 Sep, 2021 @ 11:17am"
	groups := timestampRegexp.FindStringSubmatch(timestamp)

	if len(groups) != 7 {
		return time.Time{}, false
	}

	year, _ := strconv.Atoi(groups[3])
	months := MonthStrToInt(groups[2])
	days, _ := strconv.Atoi(groups[1])
	hour, _ := strconv.Atoi(groups[4])
	if strings.ToLower(groups[6]) == "pm" {
		hour = hour + 12
	}
	mins, _ := strconv.Atoi(groups[5])

	return time.Date(
		year,
		months,
		days,
		hour,
		mins,
		0,
		0,
		time.FixedZone("PT", -7),
	), true
}

func MonthStrToInt(month string) time.Month {
	switch strings.ToLower(month) {
	case "jan", "january":
		return time.January
	case "feb", "february":
		return time.February
	case "mar", "march":
		return time.March
	case "apr", "april":
		return time.April
	case "may":
		return time.May
	case "jun", "june":
		return time.June
	case "jul", "july":
		return time.July
	case "aug", "august":
		return time.August
	case "sep", "september":
		return time.September
	case "oct", "october":
		return time.October
	case "nov", "november":
		return time.November
	case "dec", "december":
		return time.December
	default:
		return -1
	}
}
