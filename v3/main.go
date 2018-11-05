package main

import "./mhcounter"

func main() {
	c := mhcounter.New()
	c.Add(3)
	c.Add(2)
	c.Add(1)
	c.MinuteCount()
	c.HourCount()
}
