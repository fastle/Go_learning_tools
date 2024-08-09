package main

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) { // 这个是复合Writer接口的Write方法的
	*c += ByteCounter(len(p))
	return len(p), nil
}