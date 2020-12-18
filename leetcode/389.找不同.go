package leetcode

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 *
 * https://leetcode-cn.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (64.06%)
 * Likes:    200
 * Dislikes: 0
 * Total Accepted:    65.1K
 * Total Submissions: 94.8K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 * 给定两个字符串 s 和 t，它们只包含小写字母。
 *
 * 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *
 * 请找出在 t 中被添加的字母。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abcd", t = "abcde"
 * 输出："e"
 * 解释：'e' 是那个被添加的字母。
 *
 *
 * 示例 2：
 *
 * 输入：s = "", t = "y"
 * 输出："y"
 *
 *
 * 示例 3：
 *
 * 输入：s = "a", t = "aa"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 * 输入：s = "ae", t = "aea"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 1000
 * t.length == s.length + 1
 * s 和 t 只包含小写字母
 *
 *
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var num1, num2 uint8
	for _, v := range s {
		num1 += uint8(v)
	}
	for _, v := range t {
		num2 += uint8(v)
	}
	return num2 - num1
}

// @lc code=end
