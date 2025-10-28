package main

import (
	"fmt"
	"math/rand"
	"time"
)

// partition 函数：将数组 arr[left...right] 分区。
// 返回基准值最终所在的位置。
func partition(arr []int, left, right int) int {
	// 随机选择基准值，并将其移到最右边
	// 随机化是为了保证平均 O(N) 的性能，避免最坏情况 O(N^2)
	rand.Seed(time.Now().UnixNano())
	pivotIndex := left + rand.Intn(right-left+1)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	pivot := arr[right] // 选择最右边的元素作为基准值
	i := left           // i 指向小于基准值的区域的下一个位置

	// 遍历数组，将小于基准值的元素移到左边
	for j := left; j < right; j++ {
		// 注意：这里是寻找第 K 大的元素，所以我们关注的是大于 pivot 的元素
		if arr[j] >= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// 最后将基准值移到正确的位置
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// findKthLargestQuickSelect 使用快速选择算法查找第 K 大的元素。
// 算法复杂度：平均 O(N)，最坏 O(N^2)。
func findKthLargestQuickSelect(nums []int, k int) int {
	n := len(nums)
	if k <= 0 || k > n {
		panic("K 值无效")
	}

	// 注意：我们寻找的是排序后索引为 (k-1) 的元素，但从数组末尾开始算。
	// 目标索引 targetIndex 对应于第 K 大的元素在降序排序后的位置。
	targetIndex := k - 1

	left, right := 0, n-1
	for left <= right {
		// 执行一次分区操作
		pivotPosition := partition(nums, left, right)

		if pivotPosition == targetIndex {
			// 如果基准值的位置恰好是目标位置，找到了
			return nums[pivotPosition]
		} else if pivotPosition < targetIndex {
			// 如果基准值位置太靠前，目标在右侧（更大的索引）
			left = pivotPosition + 1
		} else {
			// 如果基准值位置太靠后，目标在左侧（更小的索引）
			right = pivotPosition - 1
		}
	}
	return -1 // 理论上不会执行到这里
}

func main() {
	nums1 := []int{3, 2, 1, 5, 6, 4}
	k1 := 2
	result1 := findKthLargestQuickSelect(nums1, k1)
	fmt.Printf("数组 %v 中第 %d 大的值是: %d\n", nums1, k1, result1) // 预期输出: 5

	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k2 := 4
	result2 := findKthLargestQuickSelect(nums2, k2)
	fmt.Printf("数组 %v 中第 %d 大的值是: %d\n", nums2, k2, result2) // 预期输出: 4
}
