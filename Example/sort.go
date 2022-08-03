package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 8, 7, 0, 23, 45, 2, 1}
	sortArr := quickSort(arr)

	 i := Breach(sortArr, 2, 0, len(sortArr) -1)
	fmt.Printf("%v \t i = %d", sortArr, i)
}

// 归并排序
func mergerSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]

	// 分治递归
	return merge(mergerSort(left), mergerSort(right))
}

func merge(left []int, right []int) []int {

	var result []int

	// 循环结束条件为 len(left) = 0 || len(right) = 0， 然后将有剩余数据的数组拷贝至result
	for len(left) > 0 && len(right) > 0 {
		// 取出 第0个元素，判断大小
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// 如果 left 子数组有剩余数据，拷贝值 result
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	// 如果 right 子数组有剩余数据，拷贝值 result
	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, p int, r int) []int {
	if p < r {
		//如果要排序数组中下标从 p 到 r 之间的一组数据，
		//我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
		//我们遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。
		//经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的。
		partIndex := partition(arr, p, r)
		_quickSort(arr, p, partIndex-1)
		_quickSort(arr, partIndex+1, r)
	}
	return arr
}

func partition(arr []int, p int, r int) int {
	fmt.Printf("获取锚点--%d, %d, %v\n", p, r, arr)

	pivot := arr[r]
	i := p

	for j := p; j < r; j++ {
		if arr[j] < pivot {
			swap(arr, i, j)
			i += 1
		}
	}
	//fmt.Printf("i=%d, arr=%v\n", i, arr)

	// 经过上边的循环 arr 中 i-1 的元素 小于 pivot， 将 pivot 和 arr[i] 转换位置，就变成了 arr[i-1] < arr[i] < arr[i+1] 到 arr[r]
	swap(arr, i, r)

	fmt.Printf("i=%d, arr=%v\n", i, arr)

	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func Breach(arr []int, s int, left int, right int) int {

	if left > right {
		return -1
	}

	mid := (left + right) / 2
	fmt.Printf("--%d\n", arr[mid])
	if arr[mid] == s {
		return mid
	} else if arr[mid] > s {
		end := mid - 1
		return Breach(arr, s, left, end)
	} else {
		start := mid + 1
		return Breach(arr, s, start, right)
	}
}