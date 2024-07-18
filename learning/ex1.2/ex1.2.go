package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, s := range os.Args[1:] {
		fmt.Println(idx, s)
	}
}
