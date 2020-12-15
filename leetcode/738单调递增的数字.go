package leetcode

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 *
 * https://leetcode-cn.com/problems/monotone-increasing-digits/description/
 *
 * algorithms
 * Medium (44.21%)
 * Likes:    92
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 20.2K
 * Testcase Example:  '10'
 *
 * 给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
 *
 * （当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
 *
 * 示例 1:
 *
 * 输入: N = 10
 * 输出: 9
 *
 *
 * 示例 2:
 *
 * 输入: N = 1234
 * 输出: 1234
 *
 *
 * 示例 3:
 *
 * 输入: N = 332
 * 输出: 299
 *
 *
 * 说明: N 是在 [0, 10^9] 范围内的一个整数。
 *
 */

// @lc code=start
func monotoneIncreasingDigits(N int) int {
	// 要求每一个高位都比低位小，则，如果低位小于高位，需要把低位置为9，高位减一
	// 1565-> 1559

	i := 1
	res := N
	for i <= res/10 {
		n := res / i % 100
		i *= 10
		// 高位大于低位
		if n/10 > n%10 { // 先除再乘，去除了不符合要求的两位数，保持整数量级不变
			res = res/i*i - 1
		}
	}
	return res
}

// @lc code=end
