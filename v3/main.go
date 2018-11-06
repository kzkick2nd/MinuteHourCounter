package main

import (
	"fmt"

	"./mhcounter"
)

func main() {
	c := mhcounter.NewCounter()
	c.Add(3)
	c.Add(2)
	c.Add(1)
	fmt.Println(c.MinuteCount())
	fmt.Println(c.HourCount())
}
