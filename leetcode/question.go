package leetcode

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
链接：https://leetcode-cn.com/problems/find-pivot-index
*/

/**/
func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	tmp := 0
	for j := 0; j < len(nums); j++ {

		if tmp == sum-nums[j]-tmp {
			return j
		}
		tmp += nums[j]
	}

	return -1
}

func pivotIndex2(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	tmp := 0
	cnt := len(nums)
	for j := 0; j < cnt; j++ {

		if tmp*2+nums[j] == sum {
			return j
		}
		tmp += nums[j]
	}

	return -1
}

/*
加油站
在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。
说明:
	如果题目有解，该答案即为唯一答案。
	输入数组均为非空数组，且长度相同。
	输入数组中的元素均为非负数。
链接：https://leetcode-cn.com/problems/gas-station

示例：
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

输出: 3
*/
func canCompleteCircuit(gas []int, cost []int) int {
	// 起点
	start := 0
	// 跑完后剩余油量
	run := 0
	// 剩余油量
	rest := 0

	for i := range gas {
		run += (gas[i] - cost[i])
		rest += (gas[i] - cost[i])
		if run < 0 {
			start = i + 1
			run = 0
		}
	}

	if rest >= 0 {
		return start
	}
	return -1
}

/*
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

说明:
	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。
*/
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return
}