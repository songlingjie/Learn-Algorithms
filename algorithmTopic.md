##### 斐波那耶数列

```go
// FbByRecursion 递归实现取斐波那耶数列第n个值
func FbByRecursion(n int) int {
	if n < 2 {
		return n
	}

	return FbByRecursion(n-1) + FbByRecursion(n-2)
}

// FbByArr 使用数组循环实现取斐波那耶数列第n个值
func FbByArr(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 1; i < n; i++ {
		tmp := arr[i]
		arr[i+1] = arr[i-1] + tmp
	}

	return arr[n-1]
}
```

##### 无序数组第K大元素

```go
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

```
##### 二分查找

```go
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
```
