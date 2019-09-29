package leetcode

import (
	"container/list"
)

/**
递归反转字符串
*/
func reverse(s []byte) []byte {
	if s == nil {
		return s
	}
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return s
	}
	s[0], s[len(s)-1] = s[len(s)-1], s[0]
	reverse(s[1 : len(s)-1])
	return s
}

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/
/**
时间复杂度 O(n^2)
*/
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		tem := target - v
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == tem {
				return []int{i, j}
			}
		}
	}
	return nil
}

/**
时间复杂度 O(n)，以空间换时间
临界条件判断, 差值等于当前值时不符合要求
 */
func twoSum2(nums []int, target int) []int {
	var tempMap = make(map[int]int, len(nums))
	for i, v := range nums {
		tempMap[v] = i
	}
	for i, val := range nums {
		temp := target - val
		idx, ok := tempMap[temp];
		if ok {
			if idx == i {
				continue
			}
			return []int{i, idx}
		}
	}
	return nil
}

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
  输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
  输出：7 -> 0 -> 8
  原因：342 + 465 = 807
 */

func addTwoNumbers(l1 *list.List, l2 *list.List) *list.List{
	// todo list

	return nil
}

func lengthOfLongestSubstring(s string) int {

	// 1。 将字符串放入一个map里面，

	return 0
}