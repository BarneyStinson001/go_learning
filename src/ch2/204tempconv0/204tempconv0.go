package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const AbsoluteZeroC Celsius = -273.15
const FreezingC Celsius = 0
const BoilingC Celsius	=  100

func CTOF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FTOC(f Fahrenheit) Celsius {
	return Celsius((f - 32)* 5/9)
}

func main() {
	fmt.Printf("%g\n",BoilingC-FreezingC)
	BoilingF :=CTOF(BoilingC)
	FreezingF :=CTOF(FreezingC)
	fmt.Printf("%g - %g =%g\n",BoilingF,FreezingF,BoilingF-FreezingF)
}