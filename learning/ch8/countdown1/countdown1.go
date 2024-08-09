package main

import (
	"fmt"
	"time"
)

func launch(){
	fmt.Println("Lift off!")
}
func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown--{
		fmt.Println(countdown)
		<-tick // 等计时器滴答
	}
	launch()
}



