# GO? GO!

- 使用书籍  Go程序设计语言

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

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := " ", " "
	for _, arg := range os.Args[1:] {  // 空标识符
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
```

- strings.Join(slice, string) 以string链接各元素

```go
// 练习1.1

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

```

```go
// 练习1.2

package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, s := range os.Args[1:] {
		fmt.Println(idx, s)
	}
}
```

```go
// 练习1.3 暂时不会
```


## 基本逻辑工作

```go
// 输出标准输入中出现次数大于1的行，并统计次数

package main

import (
	"bufio" // 处理输入输出， Scanner 可以读取输入，以行或者单词为单位断开， 处理以行为单位的输入内容的最简单方式， 
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // map 将string 映射到int ， 值得注意的是， 打印map的打印结果分布是随机的，设计目的是防止程序依赖某种特殊的序列, make 是map内置的函数，多种用途
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

```

- 转义字符，Go语言一般称之为verb 比较不同的是%t bool型， %q 带引号字符串， %v 内置格式的任何值， %T 任何值的类型

## 从文件读入
```go

// 打印输入中出现多次的行的个数和文本
// 可以支持文件列表读入或者stdin读入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)  //返回两个值, 一个是 *os.File 另一个是err, nil 是内置的表示没有err
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int){ //
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++  // 这里传map 是传的副本, 在函数内做修改原来的map也会变动
	}
}
```
- 以上方法一直使用流式输入,但是也有另一种方法,直接读入一大块进内存之后再分割

```go

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n"){ // 读入后划分
			counts[line]++
		}
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

```

```go
\\ 练习1.4
// 基本逻辑, 先存起来每行的信息, 然后对于每个文件, 重复以下行动, 首先对于每个句子统计,之后判断是否重复, 之后消去影响
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) > 0{
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			for _, n := range counts{
				if n > 1{
					fmt.Printf("%s\n", arg)
					break
				}
			}
			f.Close()
			f, err = os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			decountLines(f, counts)
			f.Close()
		}
	}

}

func countLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] ++
	}
}

func decountLines(f *os.File, counts map[string]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		//print("!")
		counts[input.Text()] --
	}
}

```



## GIF 动画

```go
// 用于产生随机利萨茹图形的GIF动画
// 本地打不开, 但是联网能打开
// 联网方式, 命令后添加web选项,然后访问对应端口
package main

import (
	"image"
	"image/color"
	"image/gif" // 在导入多段路径组成的包后, 使用路径的最后一段来引用这个包
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

var palette = []color.Color{color.White, color.Black} //复合字面量 数组
const ( // 常量命名, 数字,字符串,bool
	whiteIndex = 0
	blackIndex = 1
)

func main(){
	//rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // 复合字面量, 结构体
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}


```

- 练习1.5

```go
//

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

var palette = []color.Color{color.RGBA{0x3D, 0x91, 0x40, 0xff}, color.Black} // 查找RGB颜色来设置对应颜色
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web"{
		handler := func(w http.ResponseWriter, r *http.Request){
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return 
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // 复合字面量, 结构体
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
	
```

- 练习1.6

```go
//

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
)

var palette = []color.Color{color.RGBA{0x3D, 0x91, 0x40, 0xff}, color.RGBA{0x29, 0x24, 0x21, 0xff}}
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web"{
		handler := func(w http.ResponseWriter, r *http.Request){
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return 
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){
	const(
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // 复合字面量, 结构体
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		tmp := uint8(1)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res{
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y * size + 0.5), tmp)
		}
		phase += 0.1
		tmp += 1  // 轮流出现
		tmp %= 2
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
```

## 获取一个URL
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body) // ioutil包现在是io包
		resp.Body.Close() //关闭数据流来避免资源泄漏，
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
```

```go
// 练习1.7
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy%v\n", err)
		}
		resp.Body.Close()
	}

}
```

```go
\\ 练习1.8
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url, "http://") && ! strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(0)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

```


```go
\\ 练习1.9
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resq, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(0)
		}
		fmt.Printf("%v\n", resq.StatusCode) // 状态码， StatusCode
		resq.Body.Close()
	}
}
```

## 服务器

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 使用handler函数 处理所有请求
	log.Fatal(http.ListenAndServe("localhost:8000", nil))// 先打印日志到标准输出， 调用os.exit(1), 到那时defer函数不会被调用
}

func handler(w http.ResponseWriter, r *http.Request){ // handler格式， 
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}



```


```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int 


func main() {
	http.HandleFunc("/", handler) 
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // Listen
}

/*
出现问题， 使用浏览器访问的时候会调用两次接口
原因是图标也算一次
不用浏览器访问就好啦


*/


func handler(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d ffff \n", count)
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

```

```go
// 更完整的例子， 报告接收到的消息头和表单数据
package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request){ // 前者输出， 后者输入

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil{ // 将err := r.ParseForm() 嵌入到if 判断条件前， 作用域缩小
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	
}
```
### http.Request 包含内容
- Header type Header map[string][]string
- Body 请求体
- GetBody 返回Body的新副本
- ContentLength int64 关联内容长度
- TransferEncoding []string  列出了从最外层到最内层的传输编码， 一般会被忽略，当发送或者接受请求时，会自动添加或者移除”chunked“传输编码
- Close bool 连接结束后是否关闭
- Host string 服务器主机地址
- Form url.Values 表单数据
- PostForm url.Values  也是表单
- MultipartForm *multipaort.Form  解析多部分表单
- Trailer Header 表示在请求体后添加附加头
- RemoteAddr string 
- RequestURL string
- TLS *tls.ConnectionState
- Cancel <-chan struct{} 一个可选通道
- Response * Response 此请求的重定向响应

## 控制流
- if for switch(switch 不需要加break) 可以通过加fallthrough来连接到下一级
- switch 可以允许不带对象
- switch 可以紧跟简短变量声明


## go doc 
- 可以在本地直接命令行阅读标准库的文档。
- 建议：在源文件的开头写注释， 每一个函数之前写一个说明函数行为的注释， 容易使得被godoc这样的工具检测到
# go-kit 基础服务

# Others

##### 概念完整性
- 概念的完整性，是指针对于一个领域，不仅了解该领域的所有对象，并且了解所有对象之间的关系。
- 了解所有对象之间的关系，并不是感性了解，而是理性了解，并不是将所有的信息都知道就可以了，需要达到一定的理性认识，达到一定的抽象才行。

```go


```