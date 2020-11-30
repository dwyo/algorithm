package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=767 lang=golang
 *
 * [767] 重构字符串
 *
 * https://leetcode-cn.com/problems/reorganize-string/description/
 *
 * algorithms
 * Medium (43.12%)
 * Likes:    172
 * Dislikes: 0
 * Total Accepted:    14.5K
 * Total Submissions: 32.6K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
 *
 * 若可行，输出任意可行的结果。若不可行，返回空字符串。
 *
 * 示例 1:
 *
 *
 * 输入: S = "aab"
 * 输出: "aba"
 *
 *
 * 示例 2:
 *
 *
 * 输入: S = "aaab"
 * 输出: ""
 *
 *
 * 注意:
 *
 *
 * S 只包含小写字母并且长度在[1, 500]区间内。
 *
 *
 */

// @lc code=start

func reorganizeString(S string) string {
	n := len(S)
	if n == 1 {
		return S
	}
	t := int(math.Ceil(float64(n) / 2))
	// 构造一个桶
	cnt := make([]int, 26)
	// 将字母归类存放进桶
	for _, ch := range S {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > t {
			return ""
		}
	}
	ans := make([]byte, n)
	// 偶数下标,奇数下标
	evenIdx, oddIdx, halfLen := 0, 1, n/2
	for i, c := range cnt[:] {
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < n {
			ans[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 {
			ans[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(ans)
}

// @lc code=end
