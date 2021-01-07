package leetcode

/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 *
 * https://leetcode-cn.com/problems/number-of-provinces/description/
 *
 * algorithms
 * Medium (60.36%)
 * Likes:    433
 * Dislikes: 0
 * Total Accepted:    95K
 * Total Submissions: 157.3K
 * Testcase Example:  '[[1,1,0],[1,1,0],[0,0,1]]'
 *
 *
 *
 * 有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c
 * 间接相连。
 *
 * 省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
 *
 * 给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而
 * isConnected[i][j] = 0 表示二者不直接相连。
 *
 * 返回矩阵中 省份 的数量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：isConnected = [[1,0,0],[0,1,0],[0,0,1]]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * n == isConnected.length
 * n == isConnected[i].length
 * isConnected[i][j] 为 1 或 0
 * isConnected[i][i] == 1
 * isConnected[i][j] == isConnected[j][i]
 *
 *
 *
 *
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	/**
	  isConnected 数组中，所有相连在一起的城市视为一个省，
	  如果没有相连在一起的城市，也视为一个省份
	*/
	ret := 0
	con := make([]bool, len(isConnected))
	// 深度优先搜索
	/**
	  初始化 con，全部都是 false
	*/
	var dfs func(int)
	dfs = func(from int) {
		// 进入后，说明被访问到了，则修改 con 标志为 true
		con[from] = true
		// 如果城市间有连接，则对其进行搜索，确定下一个关联城市
		for j, v := range isConnected[from] {
			if v == 1 && !con[j] {
				con[j] = true
				dfs(j)
			}
		}
	}
	for i := range isConnected {
		// 如果是false，则说明还未被关联，则进行搜索
		if !con[i] {
			ret++
			dfs(i)
		}
	}

	return ret
}

// @lc code=end
