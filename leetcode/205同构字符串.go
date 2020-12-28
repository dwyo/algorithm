package leetcode

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (47.72%)
 * Likes:    315
 * Dislikes: 0
 * Total Accepted:    82.3K
 * Total Submissions: 166.2K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *
 * 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 *
 * 示例 1:
 *
 * 输入: s = "egg", t = "add"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "foo", t = "bar"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: s = "paper", t = "title"
 * 输出: true
 *
 * 说明:
 * 你可以假设 s 和 t 具有相同的长度。
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	st2 := map[byte]byte{}
	ts2 := map[byte]byte{}
	for i := range s {
		s1, ok1 := st2[s[i]]
		t1, ok2 := ts2[t[i]]
		if ok1 && ok2 {
			if s1 != t[i] || t1 != s[i] {
				return false
			}
		}
		st2[s[i]] = t[i]
		ts2[t[i]] = s[i]
		if len(st2) != len(ts2) {
			return false
		}

	}

	return true
}

// @lc code=end
