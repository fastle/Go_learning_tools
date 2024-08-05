// 对tempFlag加入支持开尔文温度。
/*
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
    String() string
    Set(string) error
}


*/
package main

import (
	"flag"
	"fmt"
)


type Celsius float64
type Fahrenheit float64
type Kelvin float64 

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)


type celsiusFlag struct{ Celsius } //这里是别名

func (f *celsiusFlag) Set(s string) error {
    var unit string
    var value float64
    fmt.Sscanf(s, "%f%s", &value, &unit)
    switch unit {
    case "C", "°C":
        f.Celsius = Celsius(value)
        return nil
    case "F", "°F":
        f.Celsius = FToC(Fahrenheit(value))
        return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil 
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c)}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
func KToC(k Kelvin) Celsius {return Celsius(k + Kelvin(AbsoluteZeroC))}
func KToF(k Kelvin) Fahrenheit {return Fahrenheit(CToF(KToC(k)))}

func CToK(c Celsius) Kelvin {return Kelvin(c - AbsoluteZeroC)}

func FToK(f Fahrenheit) Kelvin {return CToK(FToC(f))}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
    f := celsiusFlag{value}
    flag.CommandLine.Var(&f, name, usage) // 因为这里已经实现了String() 和Set() 所以可以调用
    return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
    flag.Parse()
    fmt.Println(*temp)
}