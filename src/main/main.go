package main

import (
	"fmt"

	"tempconv"
)

func main() {
	temps := []tempconv.Celsius{
		tempconv.AbsoluteZeroC,
		tempconv.FreezingC,
		tempconv.BoilingC,
	}
	for _, c := range temps {
		fmt.Printf("%7.2f°C %7.2f°F\n", c, tempconv.CToF(c))
	}
}
