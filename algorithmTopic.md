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

##### 广度优先路径
> 广度优先搜索（Breadth-First-Search），我们平常都简称 BFS。直观地讲，它其实就是一种“地毯式”层层推进的搜索策略，即先查找离起始顶点最近的，然后是次近的，依次往外搜索。 
> visited 是用来记录已经被访问的顶点，用来避免顶点被重复访问。如果顶点 q 被访问，那相应的 visited[q]会被设置为 true。queue 用来存储已经被访问、但相连的顶点还没有被访问的顶点。因为广度优先搜索是逐层访问的，也就是说，我们只有把第 k 层的顶点都访问完成之后，才能访问第 k+1 层的顶点
```php
function bfs($graph, $j, $s): array
{

    $visited = []; //判断节点是否 被访问过
    $queue   = [];
    $path    = []; // 访问路径

    $queue[]     = $j;
    $visited[$j] = true;
    while (!is_null($q = array_shift($queue))) {
        foreach ($graph[$q] as $key => $item) {
            if ($visited[$key]) {
                continue;
            }
            $path[] = $key;
            if ($key == $s) {
                return $path;
            }
            $queue[]       = $key;
            $visited[$key] = true;
        }
    }
    return $path;
}
```

##### 深度优先
> 深度优先搜索（Depth-First-Search），简称 DFS。最直观的例子就是“走迷宫”。假设你站在迷宫的某个岔路口，然后想找到出口。你随意选择一个岔路口来走，走着走着发现走不通的时候，你就回退到上一个岔路口，重新选择一条路继续走，直到最终找到出口。这种走法就是一种深度优先搜索策略。
```php
function dfs($graph, $j, $s): array
{
    $visited = [];
    $path    = [];
    $found = false;

    _dfs($j, $s, $graph, $visited, $path, $found);
    return $path;
}
function _dfs($j, $s, $graph, &$visited, &$path, &$found)
{
    if ($j == $s) {
        $found = true;
    }
    if ($found) {
        return true;
    }
    $visited[$j] = true;

    foreach ($graph[$j] as $key => $item) {
        if (!$visited[$key]) {
            $path[] = $key;
             _dfs($key, $s, $graph, $visited, $path, $found);
        }
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
> 滑动窗口的left不断的右移，只要没有重复的字符就 ++ 
> 遇到重复字符 left = right，同时将 checked 置空
> 每次移动需要计算当前最大⻓度
```go
func lengthOfLongestSubstring(s string) int  {

	result, left, right := 0,0,0
	sLen := len(s)

	checked := [256]bool{}
	for left < sLen && right < sLen {
		if checked[s[right]] == false {
			checked[s[right]] = true
			right++
		} else {
			checked  = [256]bool{}
			left = right
		}
		result = max(result, right-left)
	}

	return  result
}

```

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
###### 最小窗口子串 
- 力扣.76
```go
func minWinSubSting(s1 string, s2 string) [2]int  {

	left := 0
	var  result [2]int

	s1Len := len(s1)
	s2Len := len(s2)

	checked := [256]bool{}
	str2Sum := 0
	for i := 0; i < s2Len; i++ {
		checked[s2[i]] = true
		str2Sum += int(s2[i])
	}
	sumFlag :=0
	for right := 0; right < s1Len; right++ {
		if !checked[s1[right]] {
			continue
		}
		if sumFlag == 0 {
			left = right
		}
		sumFlag += int(s1[right])
		if sumFlag != str2Sum {
			continue
		}

		if result[1] == 0 || result[1] - result[0] > right- left {
			result[0] = left
			result[1] = right
		}

		left = right
		sumFlag = 0
	}
	return  result
}

```

