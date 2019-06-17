package lengthconv

// convert meter to feet
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

// convert feet to meter
func FToM(f Feet) Meter { return Meter(f * 0.3048) }
