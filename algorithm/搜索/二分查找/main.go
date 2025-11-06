package main

import "fmt"

// 引入sort包只是为了方便测试，实际传入的数组必须已经是有序的

// BinarySearch 在一个有序的整数数组中查找目标值的索引。
// 如果找到，返回目标值的索引；如果未找到，返回 -1。
func BinarySearch(arr []int, target int) int {
	// 设定搜索范围的左右边界
	low := 0
	high := len(arr) - 1

	for low <= high {
		// 计算中间索引。使用 (low + high) / 2 可能在 low 和 high 都很大时溢出，
		// 推荐使用 low + (high - low) / 2 这种更安全的方式。
		mid := low + (high-low)/2

		// 检查中间元素是否为目标值
		if arr[mid] == target {
			return mid // 找到目标，返回索引
		} else if arr[mid] < target {
			// 中间元素小于目标值，目标值在右半部分，更新 low
			low = mid + 1
		} else {
			// 中间元素大于目标值，目标值在左半部分，更新 high
			high = mid - 1
		}
	}

	// 循环结束，仍未找到目标值
	return -1
}

func main() {
	// 示例 1: 目标存在
	sortedArr1 := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target1 := 23
	index1 := BinarySearch(sortedArr1, target1)
	fmt.Printf("数组: %v, 目标: %d, 索引: %d\n", sortedArr1, target1, index1) // 输出: 索引: 5

	// 示例 2: 目标不存在
	sortedArr2 := []int{10, 20, 30, 40, 50}
	target2 := 35
	index2 := BinarySearch(sortedArr2, target2)
	fmt.Printf("数组: %v, 目标: %d, 索引: %d\n", sortedArr2, target2, index2) // 输出: 索引: -1

	// 示例 3: 目标在边界
	sortedArr3 := []int{1, 3, 5}
	target3 := 1
	index3 := BinarySearch(sortedArr3, target3)
	fmt.Printf("数组: %v, 目标: %d, 索引: %d\n", sortedArr3, target3, index3) // 输出: 索引: 0
}
