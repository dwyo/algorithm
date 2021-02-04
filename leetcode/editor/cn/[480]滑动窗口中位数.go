package leetcode

import (
	"sort"
)

//中位数是有序序列最中间的那个数。如果序列的大小是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
//
// 例如： 
//
// 
// [2,3,4]，中位数是 3 
// [2,3]，中位数是 (2 + 3) / 2 = 2.5 
// 
//
// 给你一个数组 nums，有一个大小为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗
//口中元素的中位数，并输出由它们组成的数组。 
//
// 
//
// 示例： 
//
// 给出 nums = [1,3,-1,-3,5,3,6,7]，以及 k = 3。 
//
// 窗口位置                      中位数
//---------------               -----
//[1  3  -1] -3  5  3  6  7       1
// 1 [3  -1  -3] 5  3  6  7      -1
// 1  3 [-1  -3  5] 3  6  7      -1
// 1  3  -1 [-3  5  3] 6  7       3
// 1  3  -1  -3 [5  3  6] 7       5
// 1  3  -1  -3  5 [3  6  7]      6
// 
//
// 因此，返回该滑动窗口的中位数数组 [1,-1,-1,3,5,6]。 
//
// 
//
// 提示： 
//
// 
// 你可以假设 k 始终有效，即：k 始终小于输入的非空数组的元素个数。 
// 与真实值误差在 10 ^ -5 以内的答案将被视作正确答案。 
// 
// Related Topics Sliding Window 
// 👍 150 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func medianSlidingWindow(nums []int, k int) []float64 {
	if k == 1 {
		ans := make([]float64, len(nums))
		for i, num := range nums {
			ans[i] = float64(num)
		}
		return ans
	}

	n := len(nums)
	ans := make([]float64, 0)
	left, right := 0, k-1
	window := make([]int, k)

	for right < n {
		// 把数放进窗口中
		// 求 left ~ right 间的中位数
		// 移动窗口
		copy(window, nums[left:right+1])
		// 排序会超时，此方法不行
		// TODO 用堆、优先队列
		sort.Ints(window)

		ans = append(ans, getMedian(window, k))
		left++
		right++
	}
	return ans
}
func getMedian(winNums []int, k int) float64 {
	d := k / 2
	if k%2 == 0 {
		return (float64(winNums[d-1]) + float64(winNums[d])) / 2
	} else {
		return float64(winNums[d])
	}
}

//leetcode submit region end(Prohibit modification and deletion)
