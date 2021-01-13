package leetcode

import (
	"math"
)

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。 
//
// 注意：你不能在买入股票前卖出股票。 
//
// 
//
// 示例 1: 
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 
//
// 示例 2: 
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
// 
// Related Topics 数组 动态规划 
// 👍 1387 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	//n := len(prices)
	//dp := make([][2]int, n)
	dp0, dp1 := 0, math.MinInt64
	for _, price := range prices {
		//if i - 1 == -1 {
		//	dp[i][0] = 0
		//	dp[i][1] = -price
		//	continue
		//}
		//dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+price)))
		//dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(-price)))

		dp0 = getMax(dp0, dp1+price)
		dp1 = getMax(dp1, -price)
	}
	//return dp[n-1][0]
	return dp0
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
