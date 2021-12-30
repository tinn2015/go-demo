/*
 * @Author: your name
 * @Date: 2021-12-29 10:50:32
 * @LastEditTime: 2021-12-29 20:28:00
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go-demo/demo1/slice.go
 */

 /* 
	
 */


package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4}
	slice2 := slice[1:3]
	fmt.Println(slice2)
	fmt.Println(slice)
	slice2[0] = 99
	fmt.Println(slice2)
	fmt.Println(slice)

	/* 切片的复制 */
	fmt.Println("--------------------------")
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{5, 4, 3}
	slice5 := []int{11, 22, 33}
	slice6 := []int{44, 55, 66}
	copy(slice4, slice3) // 只会复制slice3的前3个元素到slice4中
	fmt.Println(slice4)  // [1 2 3]
	copy(slice3, slice4) // 只会复制slice4的3个元素到slice3的前3个位置	
	fmt.Println(slice3) // [1 2 3 4 5]
	copy(slice5, slice6)
	fmt.Println(slice5) // [44, 55, 66]

	/* 添加数据 */
	fmt.Println("-------------添加数据-------------")
	dd := []int{1,2,3,4,5}
	ff := []int{77,88}
	dd = append(dd, 99)
	dd = append(dd, ff...)
	fmt.Println("addend", dd)

	/* 切片的删除 */
	fmt.Println("-----------切片的删除---------------")
	slice7 := []int{1, 2, 3, 4, 5}
	slice8 := []int{5, 4, 3, 7, 8}
	slice7 = slice7[1:] // 删除开头一个元素
	fmt.Println("slice7[1:]", slice7)
	slice7 = slice7[3:] // 删除开头N个元素
	fmt.Println(slice7)
	slice8 = slice8[:len(slice8) - 2]
	fmt.Println("slice8[:2]", slice8)

	a := []int{1, 2, 3, 4, 5, 6}
	i := 2
	// append([1 2], 4, 5, 6)
	a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	fmt.Println(a)
	// a = append(a[:i], a[i+N:]) // 删除中间N个元素
	// fmt.Println(a)
	// a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	// fmt.Println(a)
	// a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
	// fmt.Println(a)

	/* 切片的遍历 */
	fmt.Println("-----------切片的遍历---------------")
	slice11 := []int{1,2,3,4,5}
	fmt.Println("切片长度", len(slice11))

	for index, value := range slice11 {
		fmt.Println("切片item", index, value)
	}
}