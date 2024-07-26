// 练习3.13 编写KB、MB的常量声明，然后扩展到YB。
// 拆分二进制
package main

func main() {
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
		PB = 1000 * TB
		EB = 1000 * PB
		ZB = 1000 * EB
		YB = 1000 * ZB
	)
	//fmt.Println(KB, MB, GB, TB, PB, EB, ZB,YB)
}