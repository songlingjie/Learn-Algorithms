##### 冒泡排序

- 冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作

<img src="https://static001.geekbang.org/resource/image/40/e9/4038f64f47975ab9f519e4f739e464e9.jpg?wh=1142*741" style="zoom:50%;" />

> 用PHP 实现
```php
function bubbleSort()
{
    $arr = [4, 5, 6, 3, 2, 1];
    $n   = count($arr);
    // 循环数组每一个元素,每一次循序都找出当前的最大值，放到数组末尾，然后当前总数-1
    for ($i = 0; $i < $n; $i++) {
        // 再次循环数组 相邻元素比较，大于则移位
        for ($j = 0; $j < $n - 1 - $i; $j++) {
            if ($arr[$j] > $arr[$j + 1]) {
                $tmp         = $arr[$j];
                $arr[$j]     = $arr[$j + 1];
                $arr[$j + 1] = $tmp;
            }
        }
    }

    return $arr;
}
```

##### 插入排序

- 数组中的数据分为两个区间，已排序区间和未排序区间。初始已排序区间只有一个元素，就是数组的第一个元素。插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。重复这个过程，直到未排序区间中元素为空，算法结束。

<img src="https://static001.geekbang.org/resource/image/b6/e1/b60f61ec487358ac037bf2b6974d2de1.jpg?wh=1142*699" style="zoom:50%;" />

> 用PHP实现
```php

function insertSort()
{
    $arr = [4, 5, 6, 3, 2, 1];

    $n = count($arr);
    if ($n <= 1) {
        return;
    }

    // 循环数组中的元素，$i - 1 为 已经排好序的部分
    for ($i = 0; $i < $n; ++$i) {
        // 循环已经排好序的部分 未排序部分第一个值 与已排好序的部分比较，写入到合适的位置
        for ($j = $i; $j > 0; $j--) {
            if ($arr[$j] < $arr[$j - 1]) {
                $tmp         = $arr[$j];
                $arr[$j]     = $arr[$j - 1];
                $arr[$j - 1] = $tmp;
            } else {
                break;
            }
        }
    }
}
```

##### 归并排序

<img src="https://static001.geekbang.org/resource/image/db/2b/db7f892d3355ef74da9cd64aa926dc2b.jpg?wh=1142*914" style="zoom:50%;" />

- 先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了
> 用go实现
```go
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
			append(result, left[0])
			left = left[1:]
		} else {
			append(result, right[0])
			right = right[1:]
		}
	}

	// 如果 left 子数组有剩余数据，拷贝至 result 
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	// 如果 right 子数组有剩余数据，拷贝至 result
	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}
```

##### 快速排序

- 如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间的最后一个元素作为 pivot（分区点）。遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的。
- 如果对数组 [6,11,3,9,8] 执行快速排序，第一次执行partition函数之后的结果为[6,3,8,9,11]

<img src="https://static001.geekbang.org/resource/image/08/e7/086002d67995e4769473b3f50dd96de7.jpg?wh=1142*859" style="zoom:50%;" />

> 用go 实现

```go
func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, p int, r int) []int {
	if p < r {
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
```

