package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2020, 7, 19, 10, 14, 23, 0, time.Local)
	fmt.Println(t)
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.Zone())
}
