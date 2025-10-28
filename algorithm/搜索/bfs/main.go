package main

import "fmt"

// BFS 算法入口
func SimpleBFS(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	var result []int

	// 1. 初始化队列：BFS 的核心，存储待访问的“状态”（这里是索引）
	// 在一维数组中，第一个“状态”就是起始索引 0
	queue := []int{0}

	// 使用一个 map 来模拟 visited 也可以，但对于简单线性数组，
	// 只要确保索引不重复入队即可，这里我们仅依赖循环边界控制。

	for len(queue) > 0 {
		// 2. 出队：获取当前要处理的索引
		currentIndex := queue[0]
		queue = queue[1:] // 模拟出队操作

		// 3. 处理当前元素 (Process Current Element)
		currentNum := nums[currentIndex]

		// 过滤逻辑：如果是偶数，则添加到结果中
		if currentNum%2 == 0 {
			result = append(result, currentNum)
		}

		// 4. 入队：将下一个待访问的索引加入队列
		nextIndex := currentIndex + 1
		if nextIndex < len(nums) {
			// BFS 总是将当前层级的下一级（下一个索引）加入队列
			queue = append(queue, nextIndex)
		}
	}

	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filteredBFS := SimpleBFS(data)
	fmt.Printf("原始数组: %v\n", data)
	fmt.Printf("BFS 过滤偶数结果: %v\n", filteredBFS) // 输出: [2 4 6 8 10]
}
