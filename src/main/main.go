package main

import (
	"fmt"

	"tempconv"
)

func main() {
	// Define a slice of temperatures.
	temps := []tempconv.Celsius{
		tempconv.AbsoluteZeroC,
		tempconv.FreezingC,
		tempconv.BoilingC,
	}

	// Show temperatures in Celsius and Fahrenheit/Kelvin.
	for _, c := range temps {
		fmt.Printf("%7.2f°C %7.2f°F\n", c, tempconv.CToF(c))
		fmt.Printf("%7.2f°C %7.2f°K\n", c, tempconv.CToK(c))
	}

	fmt.Println("---")

	// Show how String method is used.
	c := tempconv.FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%s\n", c) // no need to call String explicitly
	fmt.Printf("%v\n", c)
}
