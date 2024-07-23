// 导入tempconv包
package main

import (
	"fmt"
	"os"
	"strconv"

	"./learning/tempconv" // go 调用不同位置的包 ，https://blog.csdn.net/Working_hard_111/article/details/139982343
)

func main(){
	for _, arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CToF(c))
	}
}