// 使用两个map记录两个字符串字符出现次数
//

package main

import "fmt"

func main(){
	s1, s2 := "abc", "cba"
	fmt.Println(cmp(s1, s2))
}

func cmp(s1, s2 string) bool {
	if s1 == s2{
		return false 
	}
	m1, m2 := make(map[string]int), make(map[string]int)
	for i := 0; i < len(s1); i++ {
		m1[string(s1[i])]++
	}
	for i := 0; i < len(s2); i++ {
		m2[string(s2[i])]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}