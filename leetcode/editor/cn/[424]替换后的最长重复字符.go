package leetcode

//给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
// 
//
// 注意：字符串长度 和 k 不会超过 104。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "ABAB", k = 2
//输出：4
//解释：用两个'A'替换为两个'B',反之亦然。
// 
//
// 示例 2： 
//
// 
//输入：s = "AABABBA", k = 1
//输出：4
//解释：
//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
//子串 "BBBB" 有最长重复字母, 答案为 4。
// 
// Related Topics 双指针 Sliding Window 
// 👍 315 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func characterReplacement(s string, k int) int {
	cntArr := [26]int{}
	n := len(s)
	left, right := 0, 0
	maxCnt, res := 0, 0
	for i := 0; i < n; i++ {
		// 计数
		cntArr[s[i] - 'A']++
		// 找到当前最大的连续字串
		maxCnt = max(maxCnt, cntArr[s[i] - 'A'])
		// 扩大窗口,右移
		right++
		//
		if right - left > maxCnt + k {
			cntArr[s[left] - 'A']--
			left++
		}
		res = max(res, right-left)
	}

	return res
}

func max( x, y int) int{
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
