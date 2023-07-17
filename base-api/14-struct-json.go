package main

import (
	"encoding/json"
	"fmt"
)

/*
	结构体字段的可见性与JSON序列化
*/

type People struct {
	ID      int    //  对外可见
	name    string // 对外不可见， 对当前包可见
	Address string `json:"address"` //json的时候讲Address字段改为小写address, 这种方式叫做结构体tag
}

type Student struct {
	ID   int
	Name string
}

type class struct {
	title    string // 因为是小写， 所以在json包里也是不可见的
	Students []Student
}

// student 构造函数
func newStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}
func main() {
	fmt.Println("=======================分隔符: 结构体tag===========================")
	people := People{
		ID:      1,
		name:    "ben",
		Address: "杭州",
	}
	peopleStr, err := json.Marshal(people)
	fmt.Printf("%s\n", peopleStr) // {"ID":1,"address":"杭州"}, name不可见且address为小写

	fmt.Println("=======================分隔符1:===========================")
	// 创建一个班级
	class1 := class{
		title:    "小一班",
		Students: make([]Student, 0, 20),
	}

	for i := 0; i < 10; i++ {
		student := newStudent(i, fmt.Sprintf("stu%02d", i))
		class1.Students = append(class1.Students, student)
	}
	fmt.Printf("%#v\n\r", class1)

	fmt.Println("=======================分隔符2===========================")

	// json 序列化： 就是将go语言的数据 -> JSON格式的字符串
	dataStr, err := json.Marshal(class1)
	if err != nil {
		fmt.Println("json序列化失败了")
		return
	}
	fmt.Printf("%s\n", dataStr)

	fmt.Println("=======================分隔符3===========================")

	// json反序列化： 就是将JSON字符串转为go语言的数据
	// jsonStringify := `{"Title":"小一班","Students":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"}]}`
	var c2 class
	err2 := json.Unmarshal([]byte(dataStr), &c2) // 需要字节类型的数据，因为要修改c2所以传入指针
	if err2 != nil {
		fmt.Println("json反序列化失败了")
		return
	}
	fmt.Printf("%#v\n", c2)
}
