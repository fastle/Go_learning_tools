package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for t := 0; t < T; t++ {
		var a1, a2, b1, b2 int 
		fmt.Scanf("%d %d %d %d", &a1, &a2, &b1, &b2)
		ans := 0
		if a1 > b1 {ans++}
		if a1 > b2 {ans++}
		if a2 > b1 {ans++}
		if a2 > b2 {ans++}
		fmt.Println(ans)
	}
}