package weightconv

// convert meter 2 feet
func MToF(m Meter) Feet { return Feet(m * 1200 / 3937) }

// convert feet 2 meter
func FToM(f Feet) Meter { return Meter(f * 3937 / 1200) }
