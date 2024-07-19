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