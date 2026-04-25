package main

import (
	"fmt"
	"math"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	var h string
	var m string
	var s string

	if hour == 0 {
		h = "00"
	} else if hour < 10 {
		h = fmt.Sprintf("0%d", hour)
	} else {
		h = fmt.Sprintf("%d", hour)
	}

	if minute == 0 {
		m = "00"
	} else if minute < 10 {
		m = fmt.Sprintf("0%d", minute)
	} else {
		m = fmt.Sprintf("%d", minute)
	}

	if second == 0 {
		s = "00"
	} else if second < 10 {
		s = fmt.Sprintf("0%d", second)
	} else {
		s = fmt.Sprintf("%d", second)
	}

	return fmt.Sprintf("%s:%s:%s", h, m, s)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hour := int(math.Floor(float64(seconds / 3600)))
	remaining := seconds - (3600 * hour)
	minute := int(math.Floor(float64(remaining / 60)))
	remaining = seconds - (minute * 60)
	second := remaining

	return hour, minute, second
}
