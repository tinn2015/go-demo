package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	// 信道阻塞主协程， 否则会报错go程序已退出
	done := make(chan int, 0)

	// 获取js 对象
	console := js.Global().Get("console")
	console.Call("log", "form go")
	fmt.Println("Hello WebAssembly")

	// 往js对象注入方法
	js.Global().Set("getTime", js.FuncOf(getTime))
	js.Global().Set("getSum", js.FuncOf(getSum))
	<-done
}

func getTime(this js.Value, args []js.Value) interface{} {
	time := time.Now().Unix()
	fmt.Println("time", time)
	return time
}

func getSum(this js.Value, args []js.Value) interface{} {
	data := args[0]
	// arrayConstructor := js.Global().Get("Uint8Array")
	// dataJS := arrayConstructor.New(data.Length())
	// js.CopyBytesToJS(dataJS, data)
	// fmt.Println("data", dataJS)
	sum := 0
	for i := 0; i < data.Length(); i++ {
		sum += data.Index(i).Get("num").Int()
	}
	return sum
}
