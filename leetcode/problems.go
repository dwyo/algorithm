package leetcode

import (
	"math"
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
		idx, ok := tempMap[temp]
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
ListNode from leetcode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
  输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
  输出：7 -> 0 -> 8
  原因：342 + 465 = 807

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var head *ListNode
	flag := 0
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if flag != 0 {
			sum += 1
			flag = 0
		}
		if sum > 9 {
			sum %= 10
			flag = 1
		}
		if res == nil {
			res = &ListNode{Val: sum}
			head = res
		} else {
			head.Next = &ListNode{Val: sum}
			head = head.Next
		}

	}
	if flag != 0 {
		head.Next = &ListNode{Val: 1}
	}
	return res
}

/**
无重复字符的最长子串
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}
	maxLen := 1
	for i := 0; i < sLen; i++ {
		hMap := make(map[string]int)
		hMap[s[i:i+1]] = 1
		len := 1
		nextS := s[i+1:]
		for j := range nextS {
			_, ok := hMap[nextS[j:j+1]]
			if ok {
				break
			}
			hMap[nextS[j:j+1]] = 1
			len++
			if maxLen < len {
				maxLen = len
			}
		}
	}
	return maxLen
}

/**
给定一个二进制数组， 计算其中最大连续1的个数。

示例 1:

输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.

链接：https://leetcode-cn.com/problems/max-consecutive-ones
*/
func findMaxConsecutiveOnes(nums []int) int {
	var cnt float64 = 0
	var max float64 = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			cnt = 0
		}
		max = math.Max(cnt, max)
	}
	return int(max)
}

/**
在《英雄联盟》的世界中，有一个叫 “提莫” 的英雄，他的攻击可以让敌方英雄艾希（编者注：寒冰射手）进入中毒状态。现在，给出提莫对艾希的攻击时间序列和提莫攻击的中毒持续时间，你需要输出艾希的中毒状态总时长。

你可以认为提莫在给定的时间点进行攻击，并立即使艾希处于中毒状态。


示例1:

输入: [1,4], 2
输出: 4
原因: 第 1 秒初，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒末结束。
第 4 秒初，提莫再次攻击艾希，使得艾希获得另外 2 秒中毒时间。
所以最终输出 4 秒。
示例2:

输入: [1,2], 2
输出: 3
原因: 第 1 秒初，提莫开始对艾希进行攻击并使其立即中毒。中毒状态会维持 2 秒钟，直到第 2 秒末结束。
但是第 2 秒初，提莫再次攻击了已经处于中毒状态的艾希。
由于中毒状态不可叠加，提莫在第 2 秒初的这次攻击会在第 3 秒末结束。
所以最终输出 3 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/teemo-attacking
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findPoisonedDuration(timeSeries []int, duration int) int {
	length := len(timeSeries)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return duration
	}
	time := 0
	for i := 0; i < length-1; i++ {
		diff := timeSeries[i+1] - timeSeries[i]
		if diff >= duration {
			time += duration
		} else {
			time += diff
		}
	}
	return time + duration
}

/*
找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
请找出数组中任意一个重复的数字。

链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
示例 1：
	输入：
	[2, 3, 1, 0, 2, 5, 3]
	输出：2 或 3
*/
func findRepeatNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		if numMap[num] > 0 {
			return num
		}
		numMap[num] = 1
	}
	return -1
}
