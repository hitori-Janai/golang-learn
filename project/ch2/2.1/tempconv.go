// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempcon performs Celsius and Fahrenheit conversions.
package tempcon

import "fmt"

// Celsius .
type Celsius float64
// Fahrenheit float64
type Fahrenheit float64
// Kelvin float64
type Kelvin float64

const (
	// AbsoluteZeroC .
	AbsoluteZeroC Celsius = -273.15
	// FreezingC     Celsius = 0
	FreezingC     Celsius = 0
	// BoilingC      Celsius = 100
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }

//!-
