/*
 * @Author: your name
 * @Date: 2021-12-28 16:39:06
 * @LastEditTime: 2021-12-28 22:46:54
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /base-test/array.go
 */
package main

import "fmt"

type point struct {
	a, b int
}

/*
* go中的数组长度和类型是固定的
*
*
*/

func main()  {
	// 申明数组并赋值
	var a [3]string = [3]string{"aa", "bb", "cc"}

	// 数组的长度根据初始化值来计算
	b := [...]int{1,2,3}
	
	fmt.Println(b)

	fmt.Println(a[0])
	// 数组长度
	fmt.Println(len(a))

	// 遍历数组
	for i, v := range a {
		fmt.Println(i, v)
	}

	for i :=0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	
	// 比较两个数组是否相等
	// 如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，
	// 只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
	aa := [2]int{1, 2}
	bb := [...]int{1, 2}
	cc := [2]int{1, 3}
	fmt.Println(aa == bb, aa == cc, bb == cc) // "true false false"
	// dd := [3]int{1, 2}
	// fmt.Println(aa == dd) // 编译错误：无法比较 [2]int == [3]int
}