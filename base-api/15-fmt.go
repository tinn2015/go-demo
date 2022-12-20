package main

import "fmt"

type student struct {
	age  int
	name string
}

func main() {
	student1 := student{
		age:  12,
		name: "ben",
	}
	fmt.Printf("年纪是：%s\n", student1.name)
}
