package main

import "fmt"

// dfsArrayFilter 是 DFS 递归函数
// nums: 输入的数组
// index: 当前正在处理的数组索引
// result: 存储结果的切片（通过引用传递，或在返回时合并）
func dfsArrayFilter(nums []int, index int, result *[]int) {
	// 1. 终止条件 (Base Case): 索引超出数组范围
	if index >= len(nums) {
		return
	}

	// 2. 处理当前元素 (Process Current Element)
	currentNum := nums[index]

	// 过滤逻辑：如果是偶数，则添加到结果中
	if currentNum%2 == 0 {
		*result = append(*result, currentNum)
	}

	// 3. 递归调用 (Recursive Step): 深入到下一个元素
	// DFS 总是优先处理当前路径的下一个节点
	dfsArrayFilter(nums, index+1, result)

	// 注意: 如果这是回溯问题（如子集、排列），这里需要“撤销”操作。
	// 但对于简单的线性过滤，不需要撤销操作。
}

// DFS 算法入口
func SimpleDFS(nums []int) []int {
	var result []int
	dfsArrayFilter(nums, 0, &result) // 从索引 0 开始
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filteredDFS := SimpleDFS(data)
	fmt.Printf("原始数组: %v\n", data)
	fmt.Printf("DFS 过滤偶数结果: %v\n", filteredDFS) // 输出: [2 4 6 8 10]
}
