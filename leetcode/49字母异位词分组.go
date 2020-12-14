package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (64.06%)
 * Likes:    550
 * Dislikes: 0
 * Total Accepted:    128.6K
 * Total Submissions: 199.7K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 很容易想到排序每个单词，用排序后的单词做key， 原单词做值，遍历完成后，返回 key 组成的数组即可
	ret := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		ret[sortedStr] = append(ret[sortedStr], str)

	}
	ans := [][]string{}
	for _, v := range ret {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
