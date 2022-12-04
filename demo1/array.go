/*
 * @Author: your name
 * @Date: 2021-12-28 16:39:06
 * @LastEditTime: 2021-12-29 10:49:13
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

func main() {

	/*
		四种初始化方式
	*/

	// 第一种方式
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("第一种方式", arr1)

	// 第二种方式
	var arr2 = [3]int{1, 2, 3}
	fmt.Println("第二种方式", arr2)

	// 第三种方式
	var arr3 = [...]int{1, 2, 3}
	fmt.Println("第三种方式", arr3)

	// 第四种方式
	var arr4 = [...]int{2: 22, 0: 10, 1: 11}
	fmt.Println("第四种方式", arr4)

	// 申明数组并赋值
	var a [3]string = [3]string{"aa", "bb", "cc"}

	// 数组的长度根据初始化值来计算
	b := [...]int{1, 2, 3}

	fmt.Println(b)

	fmt.Println(a[0])
	// 数组长度
	fmt.Println(len(a))

	// 遍历数组
	for i, v := range a {
		fmt.Println(i, v)
	}

	for i := 0; i < len(a); i++ {
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

	/* 多维数组 */

	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	// 二维码数组赋值
	array[0][0] = 10
}
