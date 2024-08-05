package main

type a interface {
	A() string
}

type newtype struct {
	int
}

func (r newtype) A() string {
	return "a"
}

type b struct {
	c newtype
}

func Pick(now a) string {
	return ""
}

func main() {
	nowtp := newtype{1}
	nowb := b{nowtp}
	Pick(nowb)
	Pick(nowtp)
}