// 进行摄氏温度和华氏温度以及绝对温度的转换
package main


type Celsius float64
type Fahrenheit float64
type Kelvin float64 

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造

func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

func KtoC(k Kelvin) Celsius {return Celsius(k + Kelvin(AbsoluteZeroC))}

func KtoF(k Kelvin) Fahrenheit {return Fahrenheit(CToF(KtoC(k)))}

func CtoK(c Celsius) Kelvin {return Kelvin(c - AbsoluteZeroC)}

func FtoK(f Fahrenheit) Kelvin {return CtoK(FtoC(f))}