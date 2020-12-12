package leetcode

import (
	"fmt"
)

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

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：
	你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

示例 1：
	输入：nums = [5,7,7,8,8,10], target = 8
	输出：[3,4]

示例 2：
	输入：nums = [5,7,7,8,8,10], target = 6
	输出：[-1,-1]

示例 3：
	输入：nums = [], target = 0
	输出：[-1,-1]

提示：
	0 <= nums.length <= 105
	-109 <= nums[i] <= 109
	nums 是一个非递减数组
	-109 <= target <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func searchRange(nums []int, target int) []int {
	// n := len(nums)

	return nums
}

// func countPrimes(n int) int {
// 	cnt := 0
// 	f := func(num int) int {
// 		for i := 2; i*i <= num; i++ {
// 			if num%i == 0 {
// 				return 0
// 			}
// 		}
// 		return 1
// 	}

// 	for x := 2; x < n; x++ {
// 		cnt += f(x)
// 	}

// 	return cnt
// }

/**
  寻找 n 以内得质数
*/
func countPrimes(n int) int {
	d := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if !d[i] {
			cnt++
			for j := i; j < n; j += i {
				d[j] = true
			}
		}
	}
	return cnt
}

/*
861. 翻转矩阵后的得分

有一个二维矩阵 A 其中每个元素的值为 0 或 1 。
移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。
在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。
返回尽可能高的分数。

示例：
	输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
	输出：39
	解释：
	转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
	0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

提示：
	1 <= A.length <= 20
	1 <= A[0].length <= 20
	A[i][j] 是 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/score-after-flipping-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func matrixScore(A [][]int) int {
	m := len(A)
	n := len(A[0])
	// m X n 矩阵

	// 进行列变换, 反转第一行全为1
	for i := 0; i < n; i++ {
		if A[0][i] == 0 { // 反转整列
			for j := 0; j < m; j++ {
				if A[j][i] == 0 {
					A[j][i] = 1
				} else {
					A[j][i] = 0
				}
			}
		}
	}
	fmt.Println("A: ", A)
	// 对第一列进行操作，使得第一列全为1，进行行变换
	for i := 1; i < m; i++ {
		if A[i][0] == 0 {
			for j := 0; j < n; j++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}
	fmt.Println("A: ", A)
	// 计算结果
	sum := 0

	// for i := range A {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Printf("A[%d][%d]: %d, 2^%d = %.2f\n", i, j, A[i][j], n-1-j, math.Pow(2, float64(n-1-j)))
	// 	}
	// }

	// for i := range A {
	// 	for j := 0; j < n; j++ {
	// 		sum += A[i][j] * int(math.Pow(2, float64(n-1-j)))
	// 	}
	// }
	sum, cnt := 0, 0
	base := 1
	for i := n - 1; i >= 0; i-- {
		for j := range A {
			if A[j][i] == 1 {
				cnt++
			}
		}
		if cnt >= m-cnt {
			sum += cnt * base
		} else {
			sum += (m - cnt) * base
		}
		base *= 2
		cnt = 0
	}
	return sum
}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if (n < 2) {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i - 1] {
			up = down + 1
		}
		if nums[i] < nums[i - 1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}