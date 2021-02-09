package leetcode

//当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：
//
// 
// 若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]； 
// 或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。 
// 
//
// 也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。 
//
// 返回 A 的最大湍流子数组的长度。 
//
// 
//
// 示例 1： 
//
// 输入：[9,4,2,10,7,8,8,1,9]
//输出：5
//解释：(A[1] > A[2] < A[3] > A[4] < A[5])
// 
//
// 示例 2： 
//
// 输入：[4,8,12,16]
//输出：2
// 
//
// 示例 3： 
//
// 输入：[100]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= A.length <= 40000 
// 0 <= A[i] <= 10^9 
// 
// Related Topics 数组 动态规划 Sliding Window 
// 👍 94 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxTurbulenceSize(arr []int) int {
	/*
		   动态规划解法
		   f[i][0]表示以i结尾的最后为<的最长子序列
		   f[i][1]表示以i结尾的最后为>的最长子序列

		   如果当前数字比它的前一个数字大:
				f[i][0] = f[i-1][1] + 1
				f[i][1] = 1
		   如果当前数字比它的前一个数字小:
				f[i][1] = f[i-1][0] + 1
				f[i][0] = 1
		   如果当前数字与它的前一个数字相同:
				f[i][1] = f[i][0] = 1
	*/
	n := len(arr)
	dp1, dp2 := 1, 1
	ans := 1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			dp1 = dp2 + 1
			dp2 = 1
		} else if arr[i] < arr[i - 1] {
			dp2 = dp1 + 1
			dp1 = 1
		} else {
			dp1, dp2 = 1,1
		}
		ans = getmax(ans, getmax(dp1, dp2))
	}
	return ans
}

func getmax(x, y int)int{
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)