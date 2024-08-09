package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x // 每次进去一个数字， 这边就会堵塞
			
		//	fmt.Printf("make %d\n", x)
		}
		close(naturals)
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	for {
		fmt.Println(<-squares)
		time.Sleep(time.Second)
	}
}