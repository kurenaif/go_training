package weightconv

// convert kilogram to pounds
func KToP(k Kilogram) Pounds { return Pounds(k / 0.45359237) }

// convert pounds to kilogram
func PToK(f Pounds) Kilogram { return Kilogram(f * 0.45359237) }
