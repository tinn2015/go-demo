/*
 * @Author: your name
 * @Date: 2022-01-04 10:37:08
 * @LastEditTime: 2022-01-05 14:43:40
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go-demo/demo1/concurrent.go
 */

 package main

 import (
	"fmt"
	"runtime"
	"time"
 )

	func log(n int) {
		for i:=0; i<5;i++{
			fmt.Print(n)
		}
		// fmt.Println(time.Now())
	}

	func loop(n int, t int) {
		for{
			fmt.Print(n)
			time.Sleep(time.Second * time.Duration(t))
		}
	}

	func main() {
		// 获取cpu个数
		fmt.Println("cpu个数", runtime.NumCPU())

		// go log(0)
		// go log(1)

		/*
			在函数前加上go关键字表示通过协程(goroutine)的方式执行
		*/
		go loop(1, 1) // 交替输出13
		loop(3, 1)
	}