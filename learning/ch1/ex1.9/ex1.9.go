package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resq, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(0)
		}
		fmt.Printf("%v\n", resq.StatusCode) // 状态码， StatusCode
		resq.Body.Close()
	}
}