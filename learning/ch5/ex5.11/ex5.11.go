// 使用带度数的拓扑排序，最后要是有度数不为零的，就是环上的。

package main

import (
	"fmt"
	"sort"
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
	du := make(map[string]int)
	for _, v := range m {
		for _, tmp := range v{
			du[tmp] ++
		}
	}

    var order []string
    var visitAll func(now string)
    visitAll = func(now string) {
		order = append(order, now)
        items := m[now]
		for _, item := range items{
            du[item] --
            if du[item] == 0 {
			//	fmt.Println("?")
				visitAll(item)
			}
        }
    }
    var keys []string
    for key := range m {
        if du[key] == 0 {
			keys = append(keys, key)
		//	fmt.Println(key + "!")
		}
    }
    sort.Strings(keys)
    for _, key := range keys {
        visitAll(key)
    }
	for _, v := range du {
		if v != 0 {
			panic("有环")
		}
	}
    return order
}