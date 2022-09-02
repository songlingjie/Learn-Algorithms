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
		arr[i+1] = arr[i-1] + arr[i]
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
##### Tow Sum
- 力扣.1
- 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
- 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个 数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返 回结果。
```go
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
```

##### 不重复的字符长度
- 力扣.3
> 依次遍历字符串 s 中的每次字符，若不是重复字符，加入一个map中，若是重复字符，则计算长度清空Map，然后再将该字符加入此Map中。

```go
func noRepeatStrLen(s string) int {
	var sMap = make(map[int]int)
	result := 0
	for i := 0; i < len(s); i++ {
		_,ok := sMap[int(s[i])]
		if ok {
			result = max(result, len(sMap))
			sMap = map[int]int{}
		}
		sMap[int(s[i])] ++
	}
	return result
}
```
- 滑动窗口法
> 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现了重复字
符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。以此类推，
每次移动需要计算当前⻓度，并判断是否需要更新最大⻓度

##### 是否是回文数
- 力扣.9
- 正序(从左向右)和倒序(从右向左)读都是一样的整数
```go
func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}
	s := strconv.Itoa(num)
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] != s[l-1-i]{
			return false
		}
	}
	return true
}
```


