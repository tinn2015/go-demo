package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	workers   []*Worker
	taskQueue chan func()
	wg        sync.WaitGroup
}

// chan func 表示一个channel 类型

type Worker struct {
	id   int
	pool *WorkerPool
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		workers:   make([]*Worker, size),
		taskQueue: make(chan func()),
	}

	for i := 0; i < size; i++ {
		worker := &Worker{
			id:   i,
			pool: pool,
		}
		pool.workers[i] = worker
		worker.start(pool)
	}

	return pool
}

func (p *WorkerPool) AddTask(task func()) {
	p.wg.Add(1)
	// 往channel 中添加一个任务
	p.taskQueue <- task
}

/*
为结构体添加一个方法
*/
func (w *Worker) start(p *WorkerPool) {
	fmt.Println("为worker 添加了start 方法")
	go func() {
		// 从channel 中获取一个task 执行
		// 之直到这个channel被销毁， 否则这个goroutine不会被销毁
		for task := range p.taskQueue {
			// 执行这个task
			// fmt.Println("协程中收到了任务")
			task()
			w.pool.wg.Done()
		}
	}()
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
	// 关闭一个channel
	close(p.taskQueue)
}

func main() {
	pool := NewWorkerPool(5)
	for i := 0; i < 10; i++ {
		taskID := i
		pool.AddTask(func() {
			fmt.Println("Executing Task", taskID)
		})
	}

	pool.Wait()
}
