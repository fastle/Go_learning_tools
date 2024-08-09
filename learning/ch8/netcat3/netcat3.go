package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
	done:= make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //这里主要起到同步的作用， 如果
	}()
	
    mustCopy(os.Stdout, conn)
	conn.Close() // 这里关闭后可以让用户端 收到关闭通知
	<- done // 如果这里先完成的话， 会先进入堵塞等待状态
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
