package main

import "fmt"

func exchangeChange1(num1 int, num2 int) {
	num1 = 20
	num2 = 60
	fmt.Println(num1, num2)
}

func exchangeChange2(num1 *int, num2 *int) {
	*num1 = 20
	*num2 = 60
	fmt.Println(num1, num2)
}

func exchangeChange3(arr [3]string) {
	arr[0] = "aaaaaa"
	fmt.Println(arr)
}

func getNum(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	var num1 int = 10
	var num2 int = 30

	/*
		默认传参的方式，在函数执行的时候会在栈帧中复制生成一个局部变量， num1 num2
		所以函数内修改不会影响到函数外
	*/
	exchangeChange1(num1, num2)
	fmt.Println("直接传参方式", num1, num2)

	/*
		以指针作为参数， 修改的是指针指向的内存， 所以函数外部也会有改变
	*/
	exchangeChange2(&num1, &num2)
	fmt.Println("指针传参方式", num1, num2)

	/*
		引用类型作为参数
		这里与js不同， js中引用类型参数中传递的引用地址， 所以函数内修改会影响到函数外
		go中的话无论是基本类型还是引用类型，函数内都会复制这个参数。所以函数内的修改不会影响函数外。
	*/
	var a [3]string = [3]string{"aa", "bb", "cc"}
	exchangeChange3(a)
	fmt.Println("引用类型", a)

	/*
		可以对函数返回值进行命名返回
	*/
	res1, res2 := getNum(5, 2)
	fmt.Println("函数返回值命名", res1, res2)
}
