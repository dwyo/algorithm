package leetcode

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 *
 * https://leetcode-cn.com/problems/candy/description/
 *
 * algorithms
 * Hard (45.07%)
 * Likes:    362
 * Dislikes: 0
 * Total Accepted:    45.1K
 * Total Submissions: 96.4K
 * Testcase Example:  '[1,0,2]'
 *
 * 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
 *
 * 你需要按照以下要求，帮助老师给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 相邻的孩子中，评分高的孩子必须获得更多的糖果。
 *
 *
 * 那么这样下来，老师至少需要准备多少颗糖果呢？
 *
 * 示例 1:
 *
 * 输入: [1,0,2]
 * 输出: 5
 * 解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2:
 *
 * 输入: [1,2,2]
 * 输出: 4
 * 解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
 *
 */
// @lc code=start
func candy(ratings []int) int {
	// 思路：两次遍历，第一次从左遍历，如果当前孩子分数比上以为分数高，则糖果比上个孩子多1，否为发一个；
	// 第二次从右往左遍历，对于当前孩子分数高于下一个孩子分数的，且糖果数量小于等于下一个孩子的，
	// 当前糖果数量加1
	n := len(ratings)
	score := make([]int, n)
	// 先给第一个孩子发一个糖果
	score[0] = 1
	//
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			score[i] = score[i-1] + 1
		} else {
			score[i] = 1
		}
	}
	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] && score[j] <= score[j+1] {
			score[j] = score[j+1] + 1
		}
	}
	sum := 0
	for _, v := range score {
		sum += v
	}

	return sum
}

// @lc code=end
