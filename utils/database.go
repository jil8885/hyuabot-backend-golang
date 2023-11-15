package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IntervalToDuration(s string) string {
	values := strings.Split(s, " ")
	if len(values) == 3 {
		days, _ := strconv.Atoi(values[0])
		timeValue := strings.Split(values[2], ":")
		hours, _ := strconv.Atoi(timeValue[0])
		minutes, _ := strconv.Atoi(timeValue[1])
		seconds, _ := strconv.Atoi(timeValue[2])

		hours += days * 24
		if minutes > 0 {
			hours++
		}
		if seconds > 0 {
			minutes++
		}
		if days < 0 {
			return fmt.Sprintf(
				"-%02d:%02d:%02d",
				int(math.Abs(float64(hours))),
				60-minutes,
				seconds,
			)
		}
		return fmt.Sprintf(
			"%02d:%02d:%02d",
			hours,
			minutes,
			seconds,
		)
	} else if len(values) == 1 {
		timeValue := strings.Split(s, ":")
		hours, _ := strconv.Atoi(timeValue[0])
		minutes, _ := strconv.Atoi(timeValue[1])
		seconds, _ := strconv.Atoi(timeValue[2])

		return fmt.Sprintf(
			"%02d:%02d:%02d",
			hours,
			minutes,
			seconds,
		)
	}
	return "00:00:00"
}
