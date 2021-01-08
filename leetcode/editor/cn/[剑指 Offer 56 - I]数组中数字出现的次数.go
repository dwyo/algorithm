package leetcode

//一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
//
// 
//
// 示例 1： 
//
// 输入：nums = [4,1,4,6]
//输出：[1,6] 或 [6,1]
// 
//
// 示例 2： 
//
// 输入：nums = [1,2,10,4,1,4,3,3]
//输出：[2,10] 或 [10,2] 
//
// 
//
// 限制： 
//
// 
// 2 <= nums.length <= 10000 
// 
//
// 
// 👍 281 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func singleNumbers(nums []int) []int {
	n := len(nums)
	if n == 2 {
		return nums
	}
	k := 0
	for _, num := range nums {
		k ^= num
	}
	// div的初始值为1，即为0001，与 k 进行&操作，找 k 中低位为1的.
	// 如果结果为0，根据&的规则，只有2个1才是1，说明 k 的低位为0，因此左移，继续匹配，直到找到 k 对应位为1，也就是低位为1的那个。
	div := 1
	for (div & k) == 0 {
		div <<= 1
	}

	a, b := 0,0
	for _, n := range nums {
		if (div & n) != 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a, b}
}
//leetcode submit region end(Prohibit modification and deletion)
