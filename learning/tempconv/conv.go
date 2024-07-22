package tempconv

// 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造

// 华氏度转摄氏度
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}