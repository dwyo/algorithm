package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 *
 * https://leetcode-cn.com/problems/rotate-array/description/
 *
 * algorithms
 * Medium (44.18%)
 * Likes:    784
 * Dislikes: 0
 * Total Accepted:    192.8K
 * Total Submissions: 436.8K
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
 *
 * 示例 1:
 *
 * 输入: [1,2,3,4,5,6,7] 和 k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右旋转 1 步: [7,1,2,3,4,5,6]
 * 向右旋转 2 步: [6,7,1,2,3,4,5]
 * 向右旋转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 * 输入: [-1,-100,3,99] 和 k = 2
 * 输出: [3,99,-1,-100]
 * 解释:
 * 向右旋转 1 步: [99,-1,-100,3]
 * 向右旋转 2 步: [3,99,-1,-100]
 *
 * 说明:
 *
 *
 * 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
 * 要求使用空间复杂度为 O(1) 的 原地 算法。
 *
 *
 */

// @lc code=start
func rotate(nums []int, k int) {
	// 方法一，双重循环
	// n := len(nums)
	// for i := 0; i < k%n; i++ {
	// 	temp := nums[n-1]
	// 	for j := n - 1; j > 0; j-- {
	// 		nums[j] = nums[j-1]
	// 	}
	// 	nums[0] = temp
	// }

	// 方法二 翻转法
	/**
	过程模拟
	reverse "----->-->" we can get "<--<-----"
	reverse "<--" 		we can get "--><-----"
	reverse "<-----" 	we can get "-->----->"
	*/

	var reverse func(int, int)
	reverse = func(start, end int) {
		temp := nums[end]
		for j := end; j > start; j-- {
			nums[j] = nums[j-1]
			fmt.Println(nums)
		}
		nums[start] = temp
	}

	n := len(nums)
	reverse(0, n-1)
	fmt.Println(nums)
	// reverse(0, k-1)
	// fmt.Println(nums)
	// reverse(k, n-1)
	// fmt.Println(nums)
}

// @lc code=end
