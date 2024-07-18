# GO? GO!



# 第一章

## 经典helloworld
- 

```go
package main // 该声明为声明自己所属在哪个包， 而不是引用， 声明为main的为一个独立的可执行程序
import "fmt" //导入声明

func main(){ //程序入口  go语言换行敏感
    // first program
    fmt.Println("Hello,World")
    
}



```

- gofmt, 自带的格式整理，vscode配置为保存后自动运行， 但是建议直接按照要求来写
  
## 命令行参数
- os.Args, 返回是一个字符串slice s, s[0] 为命令本身

```go
// 实现将参数输出出来
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // 变量生成时会被初始化为空值
	for i := 1; i < len(os.Args); i++ { // for initialization; condition; post {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
 
```

## 声明方式
1. 短变量声明， `s := ""`, 用来表示初始化比较重要， 但是通常在一个函数内部使用，不适合包级别的变量。
2. 隐式初始化， `var s string`， 用来表示初始化不甚重要
3. 空标识符， `_, arg := range os.Args[1:]`, 当返回参数中有我们不需要的值时， 应该用空标识符下划线代替， go中不允许存在未使用的变量


# go-kit 基础服务

# Others

##### 概念完整性
- 概念的完整性，是指针对于一个领域，不仅了解该领域的所有对象，并且了解所有对象之间的关系。
- 了解所有对象之间的关系，并不是感性了解，而是理性了解，并不是将所有的信息都知道就可以了，需要达到一定的理性认识，达到一定的抽象才行。