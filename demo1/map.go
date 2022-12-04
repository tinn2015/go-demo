package main

import "fmt"

/*
	map就是键值对数据类型
*/

func main() {
	var myMap map[int]string

	// 只申明map内存是没有分配空间的
	// 必须通过make函数初始化， 才会分配空间
	myMap = make(map[int]string, 10) // 存10个键值对

	myMap[111] = "aaaa"
	myMap[222] = "bbbb"
	myMap[333] = "dddd"

	fmt.Println("myMap:", myMap)

	// 也可以直接申明复制
	myMap2 := map[int]string{444: "eeee"}
	fmt.Println("myMap2:", myMap2)

	/* map 操作 */
	userMap := map[string]interface{}{
		"name":    "tinn",
		"age":     18,
		"address": "杭州",
	}

	// 新增
	userMap["score"] = 100
	fmt.Println("userMap 新增:", userMap)

	// 修改
	userMap["score"] = 80
	fmt.Println("userMap 修改:", userMap)

	// 删除， 通过内置delete函数
	delete(userMap, "score")
	fmt.Println("userMap 删除:", userMap)

	// 清空
	// 1. go中没有提供清楚map方法，需要遍历delete
	// 2. 重新make一个新的，让原来的被gc回收
	for k, v := range userMap {
		fmt.Println("user Map遍历", k, v)
		delete(userMap, k)
	}
	fmt.Println("user Map清空", userMap)

}
