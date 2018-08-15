// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (c *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		c.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		c.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name.
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

type fahrenheitFlag struct{ Fahrenheit }

func (f *fahrenheitFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Fahrenheit = CToF(Celsius(value))
		return nil
	case "F", "°F":
		f.Fahrenheit = Fahrenheit(value)
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// FahrenheitFlag defines a Fahrenheit flag with the specified name.
func FahrenheitFlag(name string, value Fahrenheit, usage string) *Fahrenheit {
	f := fahrenheitFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Fahrenheit
}
