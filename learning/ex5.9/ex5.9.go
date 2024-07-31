// 编写函数expand，将s中的"foo"替换为f("foo")的返回值

package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "fooofffofofooffooofofofofofo"
	fmt.Printf("%s\n%s\n", s, expand(s, f))
}

func expand(s string, f func(string) string) string {
	newS := f("foo")
	return strings.Replace(s, "foo", newS, -1)
}

func f(s string) string {
	return "?" + s + "?"
}