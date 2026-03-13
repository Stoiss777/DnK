package lengthconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToM converts Feet to Meters
func FToM(f Foot) Meter { return Meter(f / 3.281) }

// MToF converts Meters to Feet
func MToF(m Meter) Foot { return Foot(m * 3.281) }
