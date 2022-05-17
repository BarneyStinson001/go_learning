package tempconv


func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FTOC(f Fahrenheit) Celsius {
	return Celsius((f - 32)* 5/9)
}


func CTOK(c Celsius) Kelvin {
	return Kelvin(c+273.15)
}

func FTOK(f Fahrenheit) Kelvin {
	return Kelvin((f - 32)* 5/9+273.15)
}

func KTOC(k Kelvin) Celsius {
	return Celsius(k-273.15)
}

func KTOF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}




