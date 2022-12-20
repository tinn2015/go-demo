package main

import "fmt"

/* 结构体的方法 */
// 也叫方法接收者

type Person struct {
	Name string
}

// 给Person结构体绑定方法sayName
func (p Person) sayName() {
	// 参数是值传递的， 这里修改不会影响外部
	p.Name = "阿飞"
	fmt.Println("sayName:", p.Name)
}

func main() {
	// 创建结构体对象
	var p Person
	p.Name = "tinn"
	p.sayName()
	fmt.Println("sayName:", p.Name)
}
