// 三个示例写一个文件里了
package main

import "strings"

// 将看起来像是 系统目录的前缀删除， 并将看起来像后缀名的部分删除
func basename (s string) string {
	for i := len(s); i >= 0; i-- {
		if s[i] == '/' {
			s = s[i + 1: ]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

// 使用strings.LastIndex 
func basename1 (s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash + 1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	} 
	return s
}

// comma 

// comma 函数用于在字符串的倒数第3位之前插入逗号，以实现英文数字读法的分隔效果。
// 例如，输入 "12345678"，输出 "1234,5678"。
// 当字符串长度小于等于3时，不再插入逗号。
func comma(s string) string {
    // 获取字符串长度
    n := len(s)
    // 如果字符串长度小于等于3，不再插入逗号，直接返回原字符串
    if n <= 3 {
        return s
    }
    // 递归调用 comma 函数，将字符串的前部分和逗号以及字符串的最后三位组合起来
    // 这里的递归调用会一直进行，直到字符串长度小于等于3为止
    return comma(s[:n - 3] + "," + s[n - 3:])
}