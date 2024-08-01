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

// topoSort 实现了对有向无环图（DAG）的节点进行拓扑排序。
// m 是一个映射，其中每个键代表图中的一个节点，对应的值是该节点指向的其他节点的列表。
// 返回值是节点的拓扑排序列表。
func topoSort(m map[string][]string) []string {
    // order 用于存储拓扑排序的结果。
    var order []string
    // seen 用于记录已经访问过的节点，以避免重复访问。
    seen := make(map[string]bool)

    // visitAll 是一个递归函数，用于遍历节点并将其按拓扑顺序添加到 order 中。
    var visitAll func(items []string)
    visitAll = func(items []string) {
        for _, item := range items {
            // 如果节点尚未被访问，则递归访问其依赖项，并将该节点添加到排序顺序中。
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                order = append(order, item)
            }
        }
    }

    // keys 用于存储 m 中所有键（节点）的列表，以便进行排序。
    var keys []string
    for key := range m {
        keys = append(keys, key)
    }

    // 对节点进行排序，以便按照一定的顺序访问它们。
    sort.Strings(keys)

    // 使用排序后的节点列表调用 visitAll，以确保节点的处理顺序符合排序结果。
    visitAll(keys)  // 这里是访问入口， 从这里开始拓扑排序

    // 返回拓扑排序结果。
    return order
}