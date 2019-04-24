package main

import "fmt"

// Tonne float64
type Tonne float64

// Kilogram float64
type Kilogram float64

func (c Tonne) String() string    { return fmt.Sprintf("%gt", c) }
func (f Kilogram) String() string { return fmt.Sprintf("%gkg", f) }

func tToK(c Tonne) Kilogram { return Kilogram(c * 1000) }

func kToT(f Kilogram) Tonne { return Tonne(f / 1000) }
