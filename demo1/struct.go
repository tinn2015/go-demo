/*
 * @Author: your name
 * @Date: 2021-12-30 14:57:10
 * @LastEditTime: 2021-12-30 15:40:03
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go-demo/demo1/struct.go
 */
package main

import "fmt"

func main() {
	fmt.Println("-----------实例结构体---------------")
	type People struct{
		name string
		age int
		sex string
	}

	man := new(People)
	man.name = "tinn"
	man.age = 18
	man.sex = "男"
	fmt.Println("man", man)
	women := &People{
		name: "akali",
		age: 12,
		sex: "女",
	}
	fmt.Println("women", women)

	fmt.Println("-----------结构体的继承---------------")
	
	type Home struct {
		address string
		city string
	}
	type Member struct {
		name string
		age int
		Home
	}

	ben := &Member{
		age: 12,
		name: "ben",
	}
	ben.address = "西湖"
	ben.city = "杭州"
	fmt.Printf("%+v\n", ben)
	fmt.Println("访问继承的类型名", ben.Home.address)
	fmt.Println("直接访问", ben.address)
}