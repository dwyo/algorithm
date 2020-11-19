import "sort"

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	maxArr := nums[length-3:]
	maxArr2 := nums[0:3]
	maxRes := 0
	res1 := maxArr[0] * maxArr[1] * maxArr[2]
	res2 := maxArr2[0] * maxArr2[1] * maxArr2[2]
	res3 := maxArr2[0] * maxArr2[1] * maxArr[2]

	maxRes = res2
	if res1 > res2 {
		maxRes = res1
	}
	if res3 > maxRes {
		maxRes = res3
	}
	return maxRes
}

// @lc code=end
