package weightconv

// convert kilogram 2 pounds
func KToP(k Kilogram) Pounds { return Pounds(k / 0.45359237) }

// convert pounds 2 kilogram
func PToK(f Pounds) Kilogram { return Kilogram(f * 0.45359237) }
