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

# 第二章
- 实体的第一个字母的大小写决定其可见性是否跨包， 如果是大写开头， 说明是导出的， 可以被自己包之外的其他程序所调用
- 包名称永远是小写纯字母
- 名称的作用域越大，就使用越长且更有意义的名称
- 驼峰式命名法，首字母缩写词往往使用相同的大小写
- go中不允许出现未被定义的变量， 所有类型的变量都应当有直接可用的零值

```go
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing %g C\n",fToC(freezingF))
	fmt.Printf("boiling %g C\n", fToC(boilingF))
}


func fToC(f float64) float64{
	return (f -32) * 5 / 9
}
```


```go
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
}

```

```go
// 第四版
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")


func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}


```
- flag包简介：https://www.cnblogs.com/sparkdev/p/10812422.html

### 类型声明
- type name underlying-type
- 一般会放在函数外面全包使用， 若首字母大写则可导出包外

```go
// 进行摄氏温度和华氏温度的转换
package main


type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
```


- 命名类型之后类似于继承，可以重新定义类型的行为， 类似于下面
```go
func (c Celsius) String() string() {return fmt.Sprintf("%g°C", c)} // fmt 在将元素输出时，会优先调用函数的toString（）方法
```


### 包
- 每个包对应一个独立的命名空间， 需要明确指出包来调用， 只有名字以大写字母开头的信息才是导出的， （汉字不导出）

- 可以将之前的代码分成两个文件， 并且导出包

```go
// 用于进行摄氏度与华氏度之间的转换   tempconv.go
package tempconv 


type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

```

```go 
package tempconv // conv.go

// 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造

// 华氏度转摄氏度
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
```

- flag 

```go
// 练习2.1 注意函数复用
// 进行摄氏温度和华氏温度以及绝对温度的转换
package main


type Celsius float64
type Fahrenheit float64
type Kelvin float64 

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}  // 构造时若两个底层是相同类型可以直接构造
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
func KtoC(k Kelvin) Celsius {return Celsius(k + Kelvin(AbsoluteZeroC))}
func KtoF(k Kelvin) Fahrenheit {return Fahrenheit(CToF(KtoC(k)))}
func CtoK(c Celsius) Kelvin {return Kelvin(c - AbsoluteZeroC)}
func FtoK(f Fahrenheit) Kelvin {return CtoK(FtoC(f))}
```


### 导入包

```go

// 导入tempconv包
package main

import (
	"fmt"
	"os"
	"strconv"

	"./learning/tempconv" // go 调用不同位置的包 ，https://blog.csdn.net/Working_hard_111/article/details/139982343
)

func main(){
	for _, arg := range os.Args[1:]{
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CToF(c))
	}
}
```

- 包的初始化， 使用init（）函数， 该函数不能被调用或者引用， 每个文件中init初始化函数在程序执行的时候直接调用
```go
// 用来统计输入数的二进制1数目
package popcount

var pc [256]byte 

func init(){
	for i := range pc { // 直接可以将slice当参数
		pc[i] = pc[i / 2] + byte(i & 1) // byte 可以返回1的个数, pc[i] 表示数字i 二进制时1的位置个数
	}
}

func PopCount(x uint64) int{
	return int(pc[byte(x >> (0 * 8))] +
		pc[byte(x >> (1 * 8))] +
		pc[byte(x >> (2 * 8))] +
		pc[byte(x >> (3 * 8))] +
		pc[byte(x >> (4 * 8))] +
		pc[byte(x >> (5 * 8))] +
		pc[byte(x >> (6 * 8))] +
		pc[byte(x >> (7 * 8))])
}
```


#### 练习2.3
- 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）

```go
// 用来统计输入数的二进制1数目
package popcount

var pc [256]byte 

func init(){
	for i := range pc { // 直接可以将slice当参数
		pc[i] = pc[i / 2] + byte(i & 1) // byte 可以返回1的个数, pc[i] 表示数字i 二进制时1的位置个数
	}
}

func PopCount(x uint64) int{

	ans := 0
	for i := 0 ; i < 8; i++ {  // 写成循环形式
		ans += int(byte(x >> (i * 8)))
	}
	return ans
}
```

#### 练习2.4 
-  用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。

```go
// 用来统计输入数的二进制1数目
package popcount

var pc [256]byte 

func init(){
	for i := range pc { // 直接可以将slice当参数
		pc[i] = pc[i / 2] + byte(i & 1) // byte 可以返回1的个数, pc[i] 表示数字i 二进制时1的位置个数
	}
}

func PopCount(x uint64) int{

	ans := 0
	for ; x != 0 ; x >>= 1{   // 每次右移一位
		if x & 1 == 1{
			ans ++
		}
	}
	return ans
}

```

#### 练习2.5
- 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。

```go
// 用来统计输入数的二进制1数目
package popcount

var pc [256]byte 

func init(){
	for i := range pc { // 直接可以将slice当参数
		pc[i] = pc[i / 2] + byte(i & 1) // byte 可以返回1的个数, pc[i] 表示数字i 二进制时1的位置个数
	}
}

func PopCount(x uint64) int{

	ans := 0
	for ; x != 0 ; x = x & (x - 1){   // x - lowbit(x)
			ans ++
	}
	return ans
}
```

### 作用域
- 作用域不等于生命周期， 作用域是编码阶段的概念，生命周期是运行时的概念
- go 中编译器会一层层地向外搜寻合适的范围， 
- for, if, switch 会产生新的词法域
- 这个部分要注意好的编码习惯， 尽量不用相同的变量名， 但是go是允许使用相同的变量名的

# 第三章

- Go语言数据类型分为四类： 基础类型、复合类型、引用类型和接口类型
- 基础类型： 数字、 字符串、 bool型、 
- 复合类型：数组、结构体、
- 引用类型： 指针、 切片、 字典、 函数、 通道
- 接口类型： 第七章

## 数据类型
- Go 在运算时要求比较严格，只允许相同类型的进行运算
- 整型分为 有符号和无符号， 每种都分为 8，16，32，64位
- 还有一种对应CPU平台的类型， int和 uint
- 还用一种无符号的整数类型uintptr, 没有具体的bit大小但是足以容纳指针
- 浮点数转整数的方式是丢弃小数部分， 然后向数轴方向折断
- 浮点数只有两种， float32 和float64
- Nan 的比较总是不成立， 但是！= 会成立
- 浮点数输出可以有%e 科学计数法， %f 小数点， 两种方法， 使用%g可以自动生成

## 运算符

- 基本上和c++ 相同
- &^  为位清空操作



## 实例

```go
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)
var sin30 = math.Sin(angle) // Go的常量是在编译之前就能确定的常量
var cos30 = math.Cos(angle)



func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
	for i:= 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) { // 返回网格顶点的坐标参数
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x, y)
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
```


### 练习3.1
- 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。修改程序跳过无效的多边形。

```go
\\练习3.1 更改的代码

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
	for i:= 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Fprintf(os.Stderr, "NAN")
			} else {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

```

### 练习3.2
- 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?

```go
// 练习3.2 更改一下z轴函数即可


func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	//z := f(x, y)
	z := eggBox(x, y)
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}


func eggBox(x, y float64) float64 {
	return math.Sin(x) + math.Sin(y) / 10
}
```

### 练习3.3
- 根据高度给每个多边形上色，那样峰值部将是红色（#ff0000），谷部将是蓝色（#0000ff）。

```go
\\ 练习3.3
package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)
var sin30 = math.Sin(angle) // Go的常量是在编译之前就能确定的常量
var cos30 = math.Cos(angle)



func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
	for i:= 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i + 1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j + 1)
			dx, dy, dz := corner(i + 1, j + 1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Fprintf(os.Stderr, "NAN")
			} else {
				//将z映射到一个较大范围

				fmt.Printf("<polygon style='fill: ")
				
				avgz := int((az + bz + cz + dz) * 10.0 + 8.0) * 18
				
				redv, bluev := 0, 0 
				if avgz <= 255 {
					redv = 0
					bluev = 255 - avgz
				} else {
					redv = avgz - 255
					bluev = 0
				}
				if redv > 255 {
					redv = 255
				}
				if bluev > 255{
					bluev = 255
				}
				
				fmt.Printf("#%02X00", redv)
				fmt.Printf("%02X", bluev)	
				fmt.Printf("' points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax, ay, bx, by, cx, cy, dx, dy)
				
			}
		}
	}
	fmt.Println("</svg>")
}
$$
func corner(i, j int) (float64, float64, float64) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) 
    return math.Sin(r) / r
}
```
### 练习3.4
- 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。
- 服务器必须设置Content-Type头部： `w.Header().Set("Content-Type", "image/svg+xml")`

```go
// 直接返回给浏览器
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)
var sin30 = math.Sin(angle) // Go的常量是在编译之前就能确定的常量
var cos30 = math.Cos(angle)



func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		getXML(w)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getXML(out io.Writer){
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	"width='%d' height='%d'>", width, height)
	for i:= 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i + 1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j + 1)
			dx, dy := corner(i + 1, j + 1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Fprintf(os.Stderr, "NAN")
			} else {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i) / cells - 0.5)
	y := xyrange * (float64(j) / cells - 0.5)
	z := f(x, y)
	sx := width / 2 + (x - y) * cos30 * xyscale
	sy := height / 2 + (x + y) * sin30 * xyscale - z * zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
```


## 复数
- 附属包含complex64 和 complex128, 注意分别对应的是float32 和 float64 是两倍的关系

- Mandelbrot图像， 对每个点进行$z_{k+1} = z^2_k + c$迭代测试， 迭代次数越多出范围的颜色越深形成的图形

```go
\\ web示例
// png格式的mandelbrot 图像
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin 
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15
	var v complex128 
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast * n}
		}
	}
	return color.Black
}

```

### 练习3.5
- 实现一个彩色的Mandelbrot图像，使用image.NewRGBA创建图像，使用color.RGBA或color.YCbCr生成颜色。

```go
// 随便调调参， 颜色还挺好看
// 练习3.5 实现彩色效果
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin 
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15
	var v complex128 
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{50, 100, 100 + n, 255 - contrast * n} // 生成RGB效果
		}
	}
	return color.RGBA{200, 200, 100, 0}
}

```

### 练习3.6
- 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分成四个子像素，实现它。
- 升采样技术， 这里要求分成四个像素， 已知本来像素的中心点和宽度， 计算其他四个中心点其实不难， 结果取个平均值

```go

// 练习3.6
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin 

			xn := []float64{x - (xmax - xmin) / width / 4, x + (xmax - xmin) / width / 4}
			yn := []float64{y - (ymax - ymin) / height / 4, y + (ymax - ymin) / height / 4}
			var rnow, gnow, bnow, anow uint32 // 因为有相加操作， 所以要大一点
			//fmt.Fprintf(os.Stderr, "%g\n", xn[0])
			for _, xnow := range xn {
				for _, ynow := range yn {
					rtmp, gtmp, btmp, atmp := mandelbrot(complex(xnow, ynow)).RGBA()
					//fmt.Fprintf(os.Stderr, "%d\n", atmp)
					rnow += rtmp >> 8
					gnow += gtmp >> 8
					bnow += btmp >> 8
					anow += atmp >> 8
				}
			}
			rnow /= 4
			gnow /= 4
			bnow /= 4
			anow /= 4
			//fmt.Fprintf(os.Stderr, "%d\n", anow)
			img.SetRGBA(px, py, color.RGBA{uint8(rnow), uint8(gnow), uint8(bnow), uint8(anow)})
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15
	var v complex128 
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast * n}
		}
	}
	return color.Black
}

```

### 练习3.7 
- 另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如$z^4-1=0$。每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。
- $f(z) = z^4-1$, 已知幂函数为解析函数， 故 $f'(z) = 4z^3$

```go

// 练习3.7
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin 
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func newton(z complex128) color.Color{ // x_{i + 1} = x_i - f(x) / f'(x)
	const iterations = 200
	const contrast = 15
	var v complex128 
	v = z
	eps := 1e-8
	ans1, ans2, ans3, ans4 := complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)
	for n := uint8(0); n < iterations; n++ {
		v = v - f(v) / diff(v)
		if cmplx.Abs(v - ans1) < eps || cmplx.Abs(v - ans2) < eps || cmplx.Abs(v - ans3) < eps || cmplx.Abs(v - ans4) < eps {
			return color.Gray{255 - contrast * n}
		}
	}
	return color.Black
}

func f(z complex128) complex128 {
	return z * z * z * z - complex(1,0)
}

func diff(z complex128) complex128 {
	return 4 * z * z * z 
}
```


### 练习3.8
- 通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的分形：complex64、complex128、big.Float和big.Rat。（后面两种类型在math/big包声明。Float是有指定限精度的浮点数；Rat是无限精度的有理数。）它们间的性能和内存使用对比如何？当渲染图可见时缩放的级别是多少？