package main

import (
	"fmt"
	"sync"
)

func main() {
	// 使用 WaitGroup 等待两个协程执行完毕
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个无缓冲的通道用于同步
	ch := make(chan struct{})

	// 协程一：打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			fmt.Println("协程一:", i)
			// 发送信号，通知协程二可以打印了
			ch <- struct{}{}
			// 等待协程二的信号
			<-ch
		}
	}()

	// 协程二：打印偶数
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			// 等待协程一的信号
			<-ch
			fmt.Println("协程二:", i)
			// 发送信号，通知协程一可以打印了
			ch <- struct{}{}
		}
	}()

	// 等待所有协程执行结束
	wg.Wait()
}
