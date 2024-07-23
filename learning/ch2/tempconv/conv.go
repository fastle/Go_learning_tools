package tempconv

// 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造

// 华氏度转摄氏度
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

// 开尔文转摄氏度
func KtoC(k Kelvin) Celsius {return Celsius(k + Kelvin(AbsoluteZeroC))}

// 开尔文转华氏度
func KtoF(k Kelvin) Fahrenheit {return Fahrenheit(CToF(KtoC(k)))}


// 摄氏度转开尔文
func CtoK(c Celsius) Kelvin {return Kelvin(c - AbsoluteZeroC)}

// 华氏度转开尔文
func FtoK(f Fahrenheit) Kelvin {return CtoK(FtoC(f))}