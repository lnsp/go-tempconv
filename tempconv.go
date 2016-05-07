// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

// Temperature in degrees Celsius
type Celsius float64

// Temperature in degrees Fahrenheit
type Fahrenheit float64

const (
	// The lowest possible temperature in °C
	AbsoluteZero Celsius = -273.15
	// The freezing point of water
	Freezing Celsius = 0
	// The boiling point of water
	Boiling Celsius = 100
)

// Converts degrees Celsius to degrees Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// Converts degrees Fahrenheit to degrees Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
