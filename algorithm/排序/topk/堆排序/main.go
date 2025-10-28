package main

import (
	"container/heap"
	"fmt"
)

// 1. 定义一个用于实现小顶堆的结构体
type MinHeap []int

// 2. 实现 sort.Interface 和 heap.Interface 接口所需的方法

// Len 返回堆的长度
func (h MinHeap) Len() int { return len(h) }

// Less 定义了比较规则：h[i] < h[j] 表示 h[i] 优先级更高。
// 对于小顶堆，值越小，优先级越高（在堆顶）。
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap 交换元素
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push 向堆中添加元素
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 从堆中移除最小元素（堆顶）
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// findKthLargest 使用大小为 K 的最小堆查找数组中第 K 个最大的元素
func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		if minHeap.Len() < k {
			// 1. 如果堆的大小不足 K，直接插入
			heap.Push(minHeap, num)
		} else if num > (*minHeap)[0] {
			// 2. 如果堆已满 K 且当前元素比堆顶元素（最小值）大
			// 弹出堆顶，并插入当前元素，保持堆中始终是最大的 K 个元素
			heap.Pop(minHeap)
			heap.Push(minHeap, num)
		}
	}

	// 3. 循环结束后，堆顶元素就是第 K 大的值
	// 因为它是最大的 K 个元素中最小的那一个。
	if minHeap.Len() > 0 {
		return (*minHeap)[0]
	}
	return 0 // 或返回错误，取决于具体需求
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	// 第 2 大的值是 5 (排序后为: 6, 5, 4, 3, 2, 1)
	fmt.Printf("数组 %v 中第 %d 大的值是: %d\n", nums, k, result) // 输出: 5

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	result2 := findKthLargest(nums2, k2)
	// 第 4 大的值是 4 (排序后为: 6, 5, 5, 4, 3, 3, 2, 2, 1)
	fmt.Printf("数组 %v 中第 %d 大的值是: %d\n", nums2, k2, result2) // 输出: 4
}
