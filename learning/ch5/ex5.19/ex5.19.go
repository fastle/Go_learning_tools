// 练习5.19   使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (res int) {
	defer func() {
		if err := recover(); err != nil {
			res = 1
		}
	}()
	panic("panic")
}