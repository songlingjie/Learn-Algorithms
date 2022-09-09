package main

import (
	"fmt"
)

func main()  {
	//res := bfStrSearch("abcdefg", "cde")

	//fmt.Printf("%d", lengthOfLongestSubstring("abcadefga"))

	fmt.Printf("%t", minWinSubSting("abcopdefgave", "ae"))

}


// 暴力匹配算法  查找 m 是否在 n 中出现
// 模式串长度为 m，主串长度为 n，那在主串中，就会有 n-m+1 个长度为 m 的子串, 暴力地对比这 n-m+1 个子串与模式串
func bfStrSearch(n string, m string) bool {
	nLen := len(n)
	mLen := len(m)
	if nLen < mLen {
		return false
	}
	for i := 0; i < nLen-mLen+ 1; i++ {
		if n[i:i+mLen] == m {
			return true
		}
	}
	return false
}

// RK 字符串匹配
// 分别对 n-m+1 个子串 求哈希值
func rkStrSearch(m string, n string) bool  {

	mLen := len(m)
	nLen := len(n)
	if mLen < nLen {
		return false
	}
	//hashMap := make(map[int]bool)

	for i := 0; i < mLen-nLen+ 1; i++ {
		// 将 m[i:i+nLen] 转成hash 放入 hashMap 中
		// hash值计算方法： 将字符串每个字符转成int * 2n  abc = int(a) * 2n + b*2*n-1 + + c*2*n-2, 思路是 将10进制转成26进制
	}
	// 计算 n 的 hash 值 判断是否在 hashMap 中
	return false
}

// 检测括号是否是闭合的 力扣.20
func isParentheses(str string) bool {

	leftSymbol := make(map[string]int)
	leftSymbol["{"] = -1
	leftSymbol["["] = -2
	leftSymbol["("] = -3

	rightSymbol := make(map[string]int)
	rightSymbol["}"] = 1
	rightSymbol["]"] = 2
	rightSymbol[")"] = 3

	var stack []string
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		if rightSymbol[s] <= 0  {
			stack = append(stack, s)
			continue
		}
		if len(stack) <= 0 {
			return false
		}
		arrEnd := len(stack) - 1
		endSymbol := stack[arrEnd]
		if (rightSymbol[s] + leftSymbol[endSymbol]) != 0 {
			return false
		}
		stack = stack[:arrEnd]
	}

	return len(stack) == 0
}

// 字符串中不重复子串长度 用map实现 力扣.3
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

// 字符串中不重复子串长度 用滑动窗口实现 力扣.3
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

//  力扣.76
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

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}



