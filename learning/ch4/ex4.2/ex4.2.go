//编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var hashMethod = flag.Int("s", 256, "hash method default:256 other:384, 512")
func main() {
	flag.Parse()
	var s string
	fmt.Printf("输入字符串")
	fmt.Scanln(&s)
	switch *hashMethod{
		case 256:
			fmt.Printf("%x 1\n", sha256.Sum256([]byte(s)))
		case 384:
			fmt.Printf("%x 2\n", sha512.Sum384([]byte(s)))
		case 512:
			fmt.Printf("%x 3\n", sha512.Sum512([]byte(s)))
		default:
			fmt.Printf("输入错误")
	}
}