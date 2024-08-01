// 将 多个字符串数组拼接成一个字符串
package main

import (
	"fmt"
)
func main() {
	strs := []string{"a", "b", "c"}
	strs2 := []string{"d", "e", "f"}
	fmt.Println(myJoin(","))
	fmt.Println(myJoin("?", strs))
	fmt.Println(myJoin("?", strs, strs))
	fmt.Println(myJoin("?", strs, strs2, strs))

}

func myJoin(s string, elems ...[]string) string {
	if len(elems) == 0 {
		return ""
	}
	var str string
	for _, elem := range elems {
		for _, e := range elem {
			str += e + s
		}
	}
	return str[:len(str)-len(s)]
}