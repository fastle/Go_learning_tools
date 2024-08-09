package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done() // 正好每个都能减到
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSize)
		} else {
			fileSize <- entry.Size() // 传输到主逻辑增加。
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries 
}

var verbose = flag.Bool("v", false, "show progress")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0{
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	
	for _, root := range roots {
		n.Add(1) // 组的数目
		go walkDir(root, &n,  fileSizes)
	}
	go func(){
		n.Wait()
		close(fileSizes)
	}()


	var tick <-chan time.Time 
	if *verbose {
		tick = time.Tick(10 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
    for  {
        select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop //该处为关闭后返回的信息。loop 可以goto 和break
			}
			nfiles++
			nbytes += size 
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
    }
    printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes) / 1e9)
}