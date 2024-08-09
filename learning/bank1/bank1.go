package main

var deposites = make(chan int)
var balances = make(chan int)

func deposite(amount int) {

}

func teller() {

}

func main() {
	go teller()
}