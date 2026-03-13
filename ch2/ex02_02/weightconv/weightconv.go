package weightconv

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string { return fmt.Sprintf("%glbs", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }

// PToK converts Pounds to Kilograms
func PToK(p Pound) Kilogram { return Kilogram(p / 2.205) }

// KToP converts Kilograms to Pounds
func KToP(k Kilogram) Pound { return Pound(k * 2.205) }