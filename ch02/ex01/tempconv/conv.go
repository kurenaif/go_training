package tempconv

// convert Celsius 2 Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// convert Fahrenheit 2 Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// convert Kelvin 2 Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// convert Celsius 2 Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbosoluteZeroC) }

// convert Kelvin 2 Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// convert Fahrenheit 2 Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
