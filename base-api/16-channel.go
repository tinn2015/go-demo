package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		无缓冲区channel
		不存放数据， 整个channel的收发都会被阻塞
		也就是说一个一个执行。
	*/
	c := make(chan string)
	s := []string{"A", "B", "C", "D"}
	go func() {
		defer close(c)
		for _, v := range s {
			fmt.Printf("send %s\n", v)
			c <- v
			time.Sleep(1 * time.Second)
		}
	}()

	for v := range c {
		fmt.Printf("received: %s\n", v)
	}

	/*
		有缓冲区channel
		具有暂存数据的功能，当channel的缓冲区满了的时候才会阻塞
	*/
	c1 := make(chan string, 1)
	s1 := []string{"A", "B", "C", "D"}
	go func() {
		defer close(c1)
		for _, v1 := range s1 {
			fmt.Printf("send s1 %s\n", v1)
			c1 <- v1
			// time.Sleep(1 * time.Second)
		}
	}()
	// time.Sleep(5 * time.Second)
	for v1 := range c1 {
		fmt.Printf("received s1: %s\n", v1)
	}
}
