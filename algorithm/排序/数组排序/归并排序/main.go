package main

import "fmt"

// MergeSort 是用户调用的主函数，是程序的入口
func MergeSort(arr []int) {
	// 当数组为空或只有一个元素时，它本身就是有序的，无需排序
	if len(arr) <= 1 {
		return
	}
	// 调用递归的辅助函数来完成排序
	mergeSort(arr, 0, len(arr)-1)
}

// mergeSort 是递归函数，用于不断地分解数组
// 它负责对 arr[low...high] 这个范围内的元素进行排序
func mergeSort(arr []int, low, high int) {
	// 基线条件：当 low >= high 时，说明子数组最多只有一个元素，递归停止。
	if low < high {
		// 1. 分解 (Divide)
		// 找到中间点，防止 (low+high) 可能的整数溢出
		mid := low + (high-low)/2

		// 递归地对左右两个子数组进行排序
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)

		// 2. 合并 (Conquer / Merge)
		// 当左右两个子数组都已排好序后，将它们合并
		merge(arr, low, mid, high)
	}
}

// merge 函数是归并排序的核心，它负责将两个已排序的子数组合并成一个
// 两个子数组分别是 arr[low...mid] 和 arr[mid+1...high]
//
//  1. 先将左右两个已排序的子数组分别拷贝到临时的 left 和 right 切片中。
//  2. 然后直接在原始数组 arr 上进行归并操作，将 left 和 right 中的元素
//     按顺序放回 arr[low...high] 的正确位置。
//  3. 这样做的好处是指针 i 和 j 都可以从 0 开始，逻辑更清晰，
//     并且省去了最后将 tmp 复制回 arr 的循环。
func merge(arr []int, low, mid, high int) {
	// 创建并拷贝左右子数组

	// 拷贝左边子数组 arr[low...mid]
	leftSize := mid - low + 1
	left := make([]int, leftSize)
	copy(left, arr[low:mid+1])

	// 拷贝右边子数组 arr[mid+1...high]
	rightSize := high - mid
	right := make([]int, rightSize)
	copy(right, arr[mid+1:high+1])

	// 归并回原始数组 arr

	// i: 指向 left 切片的指针
	// j: 指向 right 切片的指针
	// k: 指向原始数组 arr 的写入位置（从 low 开始）
	i, j, k := 0, 0, low

	// 比较 left 和 right 的元素，将较小的放回 arr
	for i < leftSize && j < rightSize {
		// 保持稳定性
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	// 收尾工作
	// 如果 left 切片还有剩余，将其余元素全部复制回 arr
	for i < leftSize {
		arr[k] = left[i]
		i++
		k++
	}

	// 如果 right 切片还有剩余，将其余元素全部复制回 arr
	for j < rightSize {
		arr[k] = right[j]
		j++
		k++
	}

	// 注意：此时 arr[low...high] 已经完全有序，
}

func main() {
	arr1 := []int{10, 80, 30, 90, 40, 50, 70, 1}
	fmt.Println("原始数组:", arr1)
	MergeSort(arr1)
	fmt.Println("排序后:", arr1)

	arr2 := []int{5, 2, 8, 1, 9, 4, 6, 3, 7}
	fmt.Println("\n原始数组:", arr2)
	MergeSort(arr2)
	fmt.Println("排序后:", arr2)

	arr3 := []int{1, 1, 1, 5, 5, 2, 2, 0}
	fmt.Println("\n原始数组:", arr3)
	MergeSort(arr3)
	fmt.Println("排序后:", arr3)
}
