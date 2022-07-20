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

