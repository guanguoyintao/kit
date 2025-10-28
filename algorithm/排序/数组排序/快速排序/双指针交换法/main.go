package main

import "fmt"

// QuickSort (Lomuto) 是主函数入口
func QuickSort(arr []int) {
	quickSortSwap(arr, 0, len(arr)-1)
}

func quickSortSwap(arr []int, low, high int) {
	if low >= high {
		return
	}
	if low < high {
		p := lomutoPartition(arr, low, high)
		quickSortSwap(arr, low, p-1)
		quickSortSwap(arr, p+1, high)
	}
}

// lomutoPartition 使用双指针交换法 (Lomuto 方案)
func lomutoPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1 // i 是小于区域的右边界

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// 将基准值换到正确的位置 (i+1)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	// 1. 定义一个待排序的数组
	testArray := []int{10, 80, 30, 90, 40, 50, 70}

	// 2. 打印排序前的数组状态
	fmt.Println("原始数组:", testArray)

	// 3. 调用快速排序函数
	QuickSort(testArray)

	// 4. 打印排序后的结果
	fmt.Println("排序后:", testArray)

	// 测试其他情况
	testArray2 := []int{5, 2, 8, 1, 9, 4, 6, 3, 7}
	fmt.Println("\n原始数组:", testArray2)
	QuickSort(testArray2)
	fmt.Println("排序后:", testArray2)
}
