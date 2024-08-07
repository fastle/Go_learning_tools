// 贪心来想， 最好插到两个相同的字母中间+3， 其次就都是+2 了
package main

import "fmt"

func main(){
	var T int 
	fmt.Scan(&T)
	for t := 0; t < T; t++{
		var s string 
		fmt.Scan(&s)
		len := len(s)
		for i := 0; i < len - 1; i++{
			if s[i] == s[i + 1] {
				
			}
		}
	}
}