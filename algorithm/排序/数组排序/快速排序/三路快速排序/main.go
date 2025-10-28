package main

import (
	"fmt"
	"math/rand/v2"
)

func QuickSort(arr []int) {
	// 为了随机性，可以先打乱数组
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	quickSortSwap(arr, 0, len(arr)-1)
}

func quickSortSwap(arr []int, low, high int) {
	if low >= high {
		return
	}

	pivot := arr[low]
	lt := low    // 小于区域的右边界
	i := low + 1 // 当前指针
	gt := high   // 大于区域的左边界

	for i <= gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[gt], arr[i] = arr[i], arr[gt]
			gt--
		} else { // arr[i] == pivot
			i++
		}
	}

	// 递归地对小于和大于区域进行排序
	quickSortSwap(arr, low, lt-1)
	quickSortSwap(arr, gt+1, high)
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
