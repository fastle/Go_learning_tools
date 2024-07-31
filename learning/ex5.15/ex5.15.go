// 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。
package main

import "fmt"


func main(){
	fmt.Println(max(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(min())
}

func max(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func min(x ...int) int {
	if len(x) == 0 {
		return 0
	}
	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}
