/*
	方法接收者
*/

package main

import "fmt"

type Person struct {
	name string
	age  int8
}

func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// 这样Person就添加可Dream方法
// 这是一个值接收者
func (p Person) Dream() {
	fmt.Printf("%s在学go语言", p.name)
}

// 这是一个指针接收者
func (p *Person) setAge(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("tinn", 20)

	p1.Dream()
}
