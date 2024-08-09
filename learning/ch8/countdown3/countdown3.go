package main

import (
	"fmt"
	"os"
	"time"
)

func launch(){
	fmt.Println("Lift off!")
}
func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	abort := make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{} 
	}()
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown--{
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}



