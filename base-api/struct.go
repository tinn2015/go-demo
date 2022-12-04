/*
 * @Author: your name
 * @Date: 2021-12-30 14:57:10
 * @LastEditTime: 2022-01-01 11:08:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go-demo/demo1/struct.go
 */
package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Println("-----------实例结构体---------------")
	type People struct {
		name string
		age  int
		sex  string
	}

	man := new(People)
	man.name = "tinn"
	man.age = 18
	man.sex = "男"
	fmt.Println("man", man)
	women := &People{
		name: "akali",
		age:  12,
		sex:  "女",
	}
	fmt.Println("women", women)

	var p1 = People{"tinn", 13, "nan"}
	fmt.Println("字面量初始化", p1)

	fmt.Println("-----------结构体的继承---------------")

	type Home struct {
		address string
		city    string
	}
	type Member struct {
		name string
		age  int
		Home
	}

	ben := &Member{
		age:  12,
		name: "ben",
	}
	ben.address = "西湖"
	ben.city = "杭州"
	fmt.Printf("%+v\n", ben)
	fmt.Println("访问继承的类型名", ben.Home.address)
	fmt.Println("直接访问", ben.address)

	/* 结构体内存对齐 */
	fmt.Println("=========================内存对齐==========================")

	fmt.Println("CPU型号：", runtime.GOARCH)

	type M1 struct {
		a int8
		b int64
		c int32
	}

	m1 := M1{}
	fmt.Println("m1大小：", unsafe.Sizeof(m1))

	type M2 struct {
		a int8
		c int32
		b int64
	}

	m2 := M2{}
	fmt.Println("m1大小：", unsafe.Sizeof(m2))

	fmt.Println("=========================结构体转换==========================")

	type S1 struct {
		age int
	}

	type S2 struct {
		age int
	}

	var s1 S1
	var s2 S2
	s1 = S1(s2)
	fmt.Println(s1)
}
