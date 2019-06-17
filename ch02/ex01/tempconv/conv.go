package tempconv

// convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// convert Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// convert Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbosoluteZeroC) }

// convert Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// convert Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
