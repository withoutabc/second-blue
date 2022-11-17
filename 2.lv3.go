//按q那个不知道可以让用户在哪边输入
//Gin框架会一些基本语法，但不太会用

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	ch3 = make(chan int)
	ch4 = make(chan int)
	wg  = sync.WaitGroup{}
)

func main() {
	wg.Add(20)
	go channel()
	go 发射()
	go 发射()
	go 发射()
	go 瞄准()
	go 瞄准()
	go 瞄准()
	go 瞄准()
	go 瞄准()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 装弹()
	go 发出去()
	wg.Wait()
}
func channel() {
	ch4 <- 5
}
func 装弹() {
	for {
		<-ch4
		fmt.Printf("装弹->")
		time.Sleep(2e9)
		ch1 <- 1
	}
}
func 瞄准() {
	for {
		<-ch1
		fmt.Printf("瞄准->")
		time.Sleep(2e9)
		ch2 <- 2
	}
}
func 发射() {

	for {
		<-ch2
		ch3 <- 3
	}
}
func 发出去() {
	defer wg.Done()
	for {
		<-ch3
		fmt.Printf("发射！\n")
		time.Sleep(1.5e9)
		ch4 <- 4
	}
}
