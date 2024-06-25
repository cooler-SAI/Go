package main

import (
	"fmt"
	"time"
)

func main() {
	tz, err := time.LoadLocation("Europe/London")
	if err != nil {
		panic(err)
	}
	dates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 2, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 3, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 4, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 5, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 6, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 7, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 8, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 9, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 10, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 11, 1, 0, 0, 0, 0, tz),
		time.Date(2024, 12, 1, 0, 0, 0, 0, tz),
	}

	for _, date := range dates {
		fmt.Println(date.Format("2006-01-02T15:04:05-07:00"))
	}
}
