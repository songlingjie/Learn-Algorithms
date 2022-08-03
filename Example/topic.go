package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 0, 23, 45, 2, 1}

	n := findKNum(arr, 3, 0, len(arr)-1)

	fmt.Printf("%d", n)
}

// 求数组第K大元素
func findKNum(arr []int, k int, left int, right int) int {

	//取数组最后一个元素 arrEnd（并不是必须去最后一个，第一个 或者第n个都可以）
	//将数组中小于arrEnd移动到左边大于arrEnd放到数组右边
	arrEnd := arr[right]
	i := left
	for j := i; j < right; j++ {
		if arr[j] < arrEnd {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]

	fmt.Printf("%d, %v\n", i, arr)

	//排序后的数组 看一下 arrEnd 的下标 i 是否 = k
	//如果 i < k 说明结果在 arr[i] -- arr[right] 中，
	//如果 i > k 说明结果在 arr[left] -- arr[i] 中
	// k-1 是因为 下标从0开始
	if i < k-1 {
		return findKNum(arr, k, i+1, right)
	} else if i > k-1 {
		return findKNum(arr, k, left, i-1)
	} else {
		return arr[i]
	}
}
