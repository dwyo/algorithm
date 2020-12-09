package leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 *
 * https://leetcode-cn.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (62.82%)
 * Likes:    795
 * Dislikes: 0
 * Total Accepted:    179.5K
 * Total Submissions: 282.9K
 * Testcase Example:  '3\n7'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
 *
 * 问总共有多少条不同的路径？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：m = 3, n = 7
 * 输出：28
 *
 * 示例 2：
 *
 *
 * 输入：m = 3, n = 2
 * 输出：3
 * 解释：
 * 从左上角开始，总共有 3 条路径可以到达右下角。
 * 1. 向右 -> 向右 -> 向下
 * 2. 向右 -> 向下 -> 向右
 * 3. 向下 -> 向右 -> 向右
 *
 *
 * 示例 3：
 *
 *
 * 输入：m = 7, n = 3
 * 输出：28
 *
 *
 * 示例 4：
 *
 *
 * 输入：m = 3, n = 3
 * 输出：6
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 题目数据保证答案小于等于 2 * 10^9
 *
 *
 */

// @lc code=start
// func uniquePaths(m int, n int) int {
// 递推公式解法
// 	dp := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		dp[i] = make([]int, n)
// 		dp[i][0] = 1
// 	}
// 	for j := 0; j < n; j++ {
// 		dp[0][j] = 1
// 	}
// 	for i := 1; i < m; i++ {
// 		for j := 1; j < n; j++ {
// 			dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 		}
// 	}
// 	return dp[m-1][n-1]
// }

func uniquePaths(m int, n int) int {
	// 排列组合解法
	// 我们要想到达终点，需要往下走n-1步，往右走m-1步，总共需要走n+m-2步。
	// 他无论往右走还是往下走他的总的步数是不会变的。
	// 也就相当于总共要走n+m-2步，往右走m-1步总共有多少种走法，很明显这就是一个排列组合问题
	// 即，求  从 N = n+m-2 个数中取 M = m-1 个的排列组合
	// 公式为：(m+n-2)! / [(m-1)! * (n-1)!]

	// 23, 12

	N := n + m - 2
	// M := m - 1
	res := 1
	for i := 1; i < m; i++ {
		res = res * (N - (m - 1) + i) / i
	}

	return res
}

// @lc code=end
