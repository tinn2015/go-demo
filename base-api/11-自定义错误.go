package main

import (
	"errors"
	"fmt"
)

func test() (err error) {
	num1 := 10
	num2 := 0

	if num2 == 0 {
		// 通过errors包 抛出一个自定义包
		return errors.New("除数不能为0")
	} else {
		res := num1 / num2
		fmt.Println("res:", res)
		return nil
	}
}

func main() {
	err := test()
	if err != nil {
		fmt.Println("得到自定义错误：", err)
		panic(err) // 终止程序， 这是一个内置函数
	}
	fmt.Println("继续执行了。。。")
}
