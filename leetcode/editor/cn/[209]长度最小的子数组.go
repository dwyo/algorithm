package leetcode

//给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回
// 0。 
//
// 
//
// 示例： 
//
// 输入：s = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 
//
// 
//
// 进阶： 
//
// 
// 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。 
// 
// Related Topics 数组 双指针 二分查找 
// 👍 537 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	left, right := 0, 0
	sum := 0
	l := 0
	/**
	滑动窗口，先增加右边指针，扩大窗口范围，计数后满足要求，在缩小窗口大小，即增加左指针，缩小窗口
	*/
	for ; right < n; right++ {
		sum += nums[right]
		for sum >= s {
			if l == 0 {
				l = right - left + 1
			} else {
				l = getMin(l, right-left+1)
			}
			left += 1
			sum -= nums[left]
		}
	}
	return l
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)
