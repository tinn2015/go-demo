package main

import "fmt"

func add(num1 int, num2 int) int {
	/*
		go中defer是关键字。
		程序执行遇到defer时不会立即执行defer后面的逻辑，而是将defer后面的语句压入一个栈中，然后继续执行后面的语句
	*/
	defer fmt.Println("num1:", num1)
	defer fmt.Println("num2:", num2)
	var sum int = num1 + num2
	fmt.Println("sum:", sum)
	/*
		栈是先进后出
		函数执行完毕以后，会从栈中去除语句执行，按照先进后出的顺序执行
		本例先打印先sum再num2再num1
	*/
	return sum
}

func main() {
	fmt.Println(add(30, 60))
}
