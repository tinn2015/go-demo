package main

import "fmt"

/* demo1 ------------------接口赋值--------------------------- */

type IDatabaser interface {
	Connect() error
	Disconnect() error
}

type Irediser interface {
	Connect() error
}

type Redis struct {
	DBName string
}

func (redis *Redis) Connect() error {
	fmt.Println("redis.DBName", redis.DBName)
	fmt.Println("Redis Connect success")
	return nil
}

func (redis *Redis) Disconnect() error {
	fmt.Println("Redis Disconnect success")
	return nil
}

/* demo2 --------------------------------------------- */
/* demo3 --------------------------------------------- */

// 定义一个数据写入器的接口
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
// *file 表示file的类型
func (*file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	// 结构体Redis 隐式实现了接口IDatabaser
	// 实例Redis
	var redis = Redis{DBName: "teacher"}
	// 把实例赋值给接口
	var idb IDatabaser = &redis
	// 方法调用接口
	idb.Connect()
	idb.Disconnect()

	/* ------------接口第二种赋值法--------------- */
	var iredis Irediser
	iredis = idb
	iredis.Connect()

	/* ------------demo3--------------- */
	// 实例化file
	f := new(file)
	fmt.Println(*f)
	// 声明一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值f，也就是*file类型
	writer = f
	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")
}
