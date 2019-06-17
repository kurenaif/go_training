package lengthconv

// convert meter 2 feet
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// convert feet 2 meter
func FToM(f Feet) Meter { return Meter(f * 0.3048) }
