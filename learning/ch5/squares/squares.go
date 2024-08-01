package main

func main() {
	f := squares()
	for i := 0; i < 10; i++ {
		println(f())
	}
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}