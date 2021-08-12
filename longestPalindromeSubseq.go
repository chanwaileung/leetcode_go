package main

/**
516. 最长回文子序列
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。

示例 2：

输入：s = "aba"
输出：3
解释：一个可能的最长回文子序列为 "aba" 。
 */

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("aba"))
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	res := make([][]int, len(s))

	for i := n - 1; i >= 0; i-- {	//从末尾开始，即i为左边界字符
		res[i] = make([]int, n)
		res[i][i] = 1		//单个字符便是长度为1的回文
		for j := i + 1; j < n; j++ {		//从i的起始坐标+1开始，并且往后搜索，即i为右边界字符
			if s[i] == s[j] {
				res[i][j] = res[i + 1][j - 1] + 2		//相等则 子字符串 + 2
			}else {
				res[i][j] = max(res[i+1][j], res[i][j-1])		//如果不相等，则在掐头或者去尾的子串中获取最大值
			}
		}
	}
	return res[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}