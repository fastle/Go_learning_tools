// 进行摄氏温度和华氏温度的转换
package main

import (
	"flag"
	"fmt"
)


type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)


type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
    var unit string
    var value float64
    fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
    switch unit {
    case "C", "°C":
        f.Celsius = Celsius(value)
        return nil
    case "F", "°F":
        f.Celsius = FToC(Fahrenheit(value))
        return nil
    }
    return fmt.Errorf("invalid temperature %q", s)
}

func (f *celsiusFlag)String() string { return fmt.Sprintf("%g°C", f.Celsius) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
    f := celsiusFlag{value}
    flag.CommandLine.Var(&f, name, usage) /// 因为这里已经实现了String() 和Set() 所以可以调用
    return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
    flag.Parse()
    fmt.Println(*temp)
}