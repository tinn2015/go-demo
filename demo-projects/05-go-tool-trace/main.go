package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	// 设置最大CPU数为1
	runtime.GOMAXPROCS(2)

	f, _ := os.Create("myTrace.dat")
	defer f.Close()

	// 开始追踪，在追踪的时候跟踪将被缓冲并写入一个我们指定的文件中
	_ = trace.Start(f)
	defer trace.Stop()

	// 咱们自定义一个任务
	ctx, task := trace.NewTask(context.Background(), "customerTask")
	defer task.End()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 启动10个协程，模拟做任务
		go func(num string) {
			defer wg.Done()

			// 标记  num
			trace.WithRegion(ctx, "nihao", func() {
				var sum, i int64
				// 模拟执行任务
				for ; i < 500000000; i++ {
					sum += i
				}
				fmt.Println("WithRegion", num, sum)
			})

		}(fmt.Sprintf("num_%02d", i))
	}
	wg.Wait()
}
