##### 冒泡排序

- 冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，重复 n 次，就完成了 n 个数据的排序工作

<img src="https://static001.geekbang.org/resource/image/40/e9/4038f64f47975ab9f519e4f739e464e9.jpg?wh=1142*741" style="zoom:50%;" />

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

