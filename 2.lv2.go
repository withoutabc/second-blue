package main

import (
	"fmt"
	"time"
)

var (
	ch     = make(chan int, 1)
	Wallet = 0 // 一贫如洗的泡泡
)

func main() {
	for i := 0; i < 10000; i++ {
		go vPaopao50()
	}
	time.Sleep(2 * time.Second)
	fmt.Println("泡泡现在有", Wallet, "元")

}

func vPaopao50() {
	ch <- 5
	Wallet += 50
	<-ch
}
