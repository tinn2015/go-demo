package main

import (
	"fmt"
	"os"
)

/*
文件权限：
// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.

*/

func main() {
	readFile()
}

func readFile() {
	// 读取文件
	// 没有这个文件自动创建，拥有读写权限
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		b := make([]byte, 12) // 声明一个切片
		n, err := f.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b), n)
	}
}
