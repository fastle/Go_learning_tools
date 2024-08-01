// 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)

    var visitAll func(m map[string][]string, s string)
    visitAll = func(m map[string][]string, s string) {
		if seen[s] {  // 多入口， 把判断改到循环开始
			return
		}
		seen[s] = true
		order = append(order, s)
		items := m[s]
        for _, item := range items {
			//fmt.Println(s + "!" + item + "!")
            visitAll(m , item)
        }
    }
    for key := range m {
		visitAll(m, key)
    }
    return order
}