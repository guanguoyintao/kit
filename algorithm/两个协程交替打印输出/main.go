package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	// turn 用来标记当前应该由哪个协程执行
	turn := 1

	// 协程一
	go func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			mu.Lock()
			// 如果不是自己的回合，就等待
			for turn != 1 {
				cond.Wait()
			}
			fmt.Println("协程一:", i)
			// 切换回合
			turn = 2
			// 唤醒可能在等待的另一个协程
			cond.Signal()
			mu.Unlock()
		}
	}()

	// 协程二
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			mu.Lock()
			// 如果不是自己的回合，就等待
			for turn != 2 {
				cond.Wait()
			}
			fmt.Println("协程二:", i)
			// 切换回合
			turn = 1
			// 唤醒可能在等待的另一个协程
			cond.Signal()
			mu.Unlock()
		}
	}()

	wg.Wait()
}
