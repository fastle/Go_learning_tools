// 传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
// f返回的是一个新的类型， 和Printf的封装不是很一样
package main

import (
	"fmt"
	"io"
	"os"
)

// CountWriter是一个结构体，用于包装一个Writer接口，并提供一个计数器，用于记录写入的字节数。
type CountWriter struct {
    Writer io.Writer // 被包装的Writer接口。
    Count  int64     // 已写入的字节数。
}

// Write方法将内容写入底层的Writer，并更新已写入的字节数。
// 该方法首先尝试将内容写入包装的Writer中，如果写入过程中发生错误，将直接返回已写入的字节数和错误。
// 如果写入成功，则增加Count的值，以反映新写入的字节数。
func (cw *CountWriter) Write(content []byte) (int, error) {
    // 调用包装的Writer的Write方法，返回写入的字节数n和可能的错误。
    n, err := cw.Writer.Write(content)
    if err != nil {
        // 如果写入过程中发生错误，返回已写入的字节数和错误。
        return n, err
    }
    // 更新已写入的字节数。
    cw.Count += int64(len(content))
    // 返回写入的字节数和nil错误，表示写入成功。
    return n, nil
}


// CountingWriter 包装一个 io.Writer 实现，返回一个新的写入器和一个指针，指向一个 int64 类型的计数器。
// 该计数器会记录写入到包装的写入器中的字节数。
// 参数 w: 被包装的 io.Writer 实例。
// 返回值:
// - 一个包装后的 io.Writer 实例，该实例会追踪写入的字节数。
// - 一个指向 int64 类型的指针，用于访问已写入的字节数。
func CountingWriter(w  io.Writer)(io.Writer, *int64) {
    // 创建并初始化 CountWriter 实例，其中 Writer 字段被设置为传入的 w 参数。
    cw := CountWriter{Writer: w}
    // 返回 CountWriter 实例的指针和其内部计数器的指针。
    return &cw, &cw.Count
}

func main() {
	cw, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "%s", "Print somethind to the screen...")
	fmt.Println(*counter)
}