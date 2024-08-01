package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
    defer printStack()
    f(3)
}
func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) 
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
/*
输出第一部分
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3  // 发生异常、 之前延迟的defer先被调用， 然后再触发panic
panic: runtime error: integer divide by zero


printStack() 的输出为
goroutine 1 [running]:
main.printStack()
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:15 +0x2e
panic({0xb27a80?, 0xbcb9f0?})
        C:/Program Files/Go/src/runtime/panic.go:770 +0x132
main.f(0xb5b098?)
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:19 +0x118
main.f(0x1)
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:21 +0xfe
main.f(0x2)
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:21 +0xfe
main.f(0x3)
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:21 +0xfe
main.main()
        D:/bq/Go_learning_tools/learning/defer2/defer2.go:11 +0x35
*/