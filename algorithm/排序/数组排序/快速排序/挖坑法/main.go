package main

import "fmt"

// QuickSort 是主函数入口
func QuickSort(arr []int) {
	quickSortFillPit(arr, 0, len(arr)-1)
}

func quickSortFillPit(arr []int, low, high int) {
	if low < high {
		pivotIndex := fillThePitPartition(arr, low, high)
		quickSortFillPit(arr, low, pivotIndex-1)
		quickSortFillPit(arr, pivotIndex+1, high)
	}
}

// fillThePitPartition 使用挖坑法进行分区
func fillThePitPartition(arr []int, low, high int) int {
	// 1. 选定基准，挖出第一个坑
	pivot := arr[low]

	for low < high {
		// 2. 高指针从右向左，找比 pivot 小的来填坑
		for low < high && arr[high] >= pivot {
			high--
		}
		// 找到小的，填到 low 的坑里
		if low < high {
			arr[low] = arr[high]
		}

		// 3. 低指针从左向右，找比 pivot 大的来填坑
		for low < high && arr[low] <= pivot {
			low++
		}
		// 找到大的，填到 high 的坑里
		if low < high {
			arr[high] = arr[low]
		}
	}

	// 5. 基准归位
	arr[low] = pivot
	return low
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
