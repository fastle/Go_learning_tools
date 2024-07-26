// 实现非递归版本的comma函数

package main

import (
	"bytes"
	"strings"
)

func main(){
	for _, s := range []string{"1", "12", "123", "1234", "12345678.901234", "-123456789012.345", "123456789012.3456", "123456789012345.67", "12345678901234567.8", "12.34567890123456789"} {
		println(comma(s))
	}
}

// 更改后的comma 函数支持将浮点数, 小数部分是从左往右
func comma(s string) string {
	var buf bytes.Buffer
	len := len(s)
	if len > 0 && s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}
	dotPlace := strings.IndexByte(s, '.') 
	if dotPlace > 0 {
		buf.WriteString(comma1(s[:dotPlace]))
		buf.WriteByte('.')
		buf.WriteString(comma2(s[dotPlace + 1:]))
	} else {
		buf.WriteString(comma1(s))
	}
	return buf.String()
}
// 整数部分
func comma1(s string) string {
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
// 小数部分
func comma2(s string) string {
	var buf bytes.Buffer
	len := len(s)
	
	for be := 0; be < len ; be += 3{
		if be + 3 < len {
			buf.WriteString(s[be:be+3])
			buf.WriteString(",")
		} else {
			buf.WriteString(s[be:])
		}
	}
	return buf.String()
}