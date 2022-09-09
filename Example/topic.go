package main

import (
	"fmt"
	"strconv"
)

func main() {
	//arr := []int{9, 8, 7, 0, 23, 45, 2, 1}

	//n := findKNum(arr, 3, 0, len(arr)-1)
	//
	//fmt.Printf("%d", n)

	//i := twoSum(arr, 15)
	//fmt.Printf("%v", i)

	fmt.Printf("result is %d", isPalindrome(12321))
}

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	s := strconv.Itoa(num)
	l := len(s)
	for i := 0; i < l / 2; i++ {
		if s[i] != s[l-1-i]{
			return false
		}
	}
	return true
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


func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		otherNum := target - nums[i]
		otherIndex,ok := hashTable[otherNum]
		if ok {
			return []int{i,otherIndex}
		}
		hashTable[nums[i]] = i
	}

	return []int{}
}

func threeSum()  {
	
}

func Max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
