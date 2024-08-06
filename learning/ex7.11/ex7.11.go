package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item") // 获取参数方法
		price, ok := db[item]
		if ! ok{
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	case "/create": 
		item := r.URL.Query().Get("item")
		price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid price: %s\n", r.URL.Query().Get("price"))
			return
		}
		db[item] = dollars(price)
	case "/modify":
		item := r.URL.Query().Get("item")
		price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid price: %s\n", r.URL.Query().Get("price"))
			return 
		}
		db[item] = dollars(price)
	case "/delete":
		item := r.URL.Query().Get("item")
		delete(db, item) // delete 为Go内置， 按照指定的键将元素从map中删除， 若删除的键为nil或者在map中不存在， 则不进行任何操作。
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", r.URL)	
	}
}