package leetcode

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (47.83%)
 * Likes:    314
 * Dislikes: 0
 * Total Accepted:    121.6K
 * Total Submissions: 249.6K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 *
 *
 * 示例：
 *
 * s = "leetcode"
 * 返回 0
 *
 * s = "loveleetcode"
 * 返回 2
 *
 *
 *
 *
 * 提示：你可以假定该字符串只包含小写字母。
 *
 */

// @lc code=start
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	// 两次遍历，一次计数，第二次查找下标
	n := len(s)
	chNum := make([]int, 26)
	for _, v := range s {
		chNum[v-'a']++
	}
	for i := 0; i < n; i++ {
		if chNum[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
