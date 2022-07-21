<?php

// 冒泡排序
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

function insertSort()
{
    $arr = [4, 5, 6, 3, 2, 1, 0];

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

insertSort();



