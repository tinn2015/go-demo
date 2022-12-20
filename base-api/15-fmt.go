package main

import "fmt"

type student struct {
	age  int
	name string
}

func main() {
	fmt.Print("直接输出", "\n")
	fmt.Println("换行并输出")

	// 格式化输出
	fmt.Printf("姓名：%s, 年龄：%d\n", "alice", 17) // %s字符串占位， %d整形的占位

	fmt.Printf("%s %d %t\n", "字符串", 17, true) // %t 布尔类型占位

	fmt.Printf("%f\n", 3.1415) // %f 浮点类型占位

	fmt.Printf("%.1f\n", 3.1415) // 指定精度， 保留一位小数

	fmt.Printf("%6.1f\n", 3.1415) // 指定宽度和精度

	student1 := student{
		age:  12,
		name: "ben",
	}
	fmt.Printf("年纪是：%s\n", student1.name)

	// 输入， 获取用户输入值
	// var name string
	// var age int
	var name2 string
	var age2 int
	// fmt.Print("请输入姓名：")
	// fmt.Scan(&name)
	// fmt.Print("请输入年龄：")
	// fmt.Scan(&age)
	// fmt.Printf("输入的姓名是：%s,输入的年龄是%d\n", name, age)

	// 格式化输入
	fmt.Print("请再次输入：")
	fmt.Scanf("姓名 %s 年龄 %d", &name2, &age2)
	fmt.Printf("输入的姓名是：%s,输入的年龄是%d\n", name2, age2)

}
