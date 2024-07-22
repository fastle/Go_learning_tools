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