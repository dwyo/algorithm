package leetcode

import (
	"math"
)

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
//
// 
//
// 示例： 
//
// 
//输入：[1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
// 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= n <= 30,000。 
// 所给数据范围 [-10,000，10,000]。 
// 
// Related Topics 数组 
// 👍 140 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	sum := 0
	if n <= k {
		for _, v := range nums {
			sum += v
		}
		return float64(sum) / float64(k)
	}

	maxSum := math.MinInt32
	for i, v := range nums {
		if i < k {
			sum += v
			maxSum = sum
			continue
		}
		sum = sum + v - nums[i - k]
		maxSum = getmax(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}
func getmax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)


/**
[9,7,3,5,6,2],0,8,1,9 	32
9,[7,3,5,6,2,0],8,1,9	23
9,7,[3,5,6,2,0,8],1,9	24
9,7,3,[5,6,2,0,8,1],9	22
9,7,3,5,[6,2,0,8,1,9]	26

 */

// 32, 23, 24,