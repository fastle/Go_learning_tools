// 实现非递归版本的comma函数

package main

import "bytes"

func main(){
	for _, s := range []string{"1", "12", "123", "1234", "12345678901234", "123456789012345", "1234567890123456", "12345678901234567", "123456789012345678", "1234567890123456789"} {
		println(comma(s))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	len := len(s)
	mod1 := len % 3
	if mod1 == 0 && len >= 3 {
		mod1 = 3
	}
	for be := 0; be + mod1 <= len ;{
		buf.WriteString(s[be:be+mod1])
		if be + mod1 != len {
			buf.WriteString(",")
		}
		be += mod1
		mod1 = 3
	}

	return buf.String()
}