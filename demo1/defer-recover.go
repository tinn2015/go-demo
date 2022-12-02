package main

import "fmt"

/*
	用defer + recover 实现错误捕获
*/

func test() {
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("result:", res)
}

func test2() {
	// defer后面加上匿名函数调用
	defer func() {
		// 调用recover内置函数，捕获错误
		err := recover()
		// 没有捕获到错误返回零值 nil
		if err != nil {
			fmt.Println("捕获到错误！")
			fmt.Println("error:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("result:", res)
}

func main() {
	/*
		再没有错误捕获的时候test函数会报错， 导致程序运行中断
	*/
	// test()

	// test2加上了错误捕获，函数执行不会中断
	test2()
	fmt.Println("继续执行了。。。")
}
