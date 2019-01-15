// Package tempconv performs Celsius and Fahrenheit/Kelvin temparature calculations.
package tempconv

import "fmt"

// Types declarations.
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// Constants declarations.
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	ZeroC         Celsius = FreezingC
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = Kelvin(AbsoluteZeroC)
)

// Conversion functions.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func KtoC(k Kelvin) Celsius     { return Celsius(k - ZeroK) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + AbsoluteZeroC) }

// String method controls how values are printed as a string byt the fmt
// package.
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
