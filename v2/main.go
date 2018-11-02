package main

import (
	"fmt"

	"./mhcounter"
)

func main() {
	var e mhcounter.Events
	e.Add(1)
	fmt.Println(e.MinuteCount())
	fmt.Println(e.HourCount())
}
