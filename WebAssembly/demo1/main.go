package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("from go")
	fmt.Println("Hello WebAssembly")
	fmt.Println("js", js.Global())
}

func getTime() {
	fmt.Println("time:", time.Now())
}
