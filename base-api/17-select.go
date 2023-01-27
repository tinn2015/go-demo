package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		select 基本使用
		没有default且没有匹配的case,那么Select就会一直阻塞
	*/
	wg := sync.WaitGroup{}
	wg.Add(1)
	f1 := make(chan struct{})
	f2 := make(chan struct{})

	go func() {
		defer wg.Done()
		select {
		case <-f1:
			fmt.Println("接收一号工厂的产品")
		case <-f2:
			fmt.Println("接收二号工厂的产品")
		}
	}()

	time.Sleep(3 * time.Second)
	f1 <- struct{}{}
	wg.Wait()

	fmt.Println("================================分割线====================================")

	/*
		通常用for + Select 来完成对channel的持续监听
	*/
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

loop:
	for {
		select {
		case i, ok := <-c:
			if !ok {
				break loop
			}
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}
