package leetcode

import (
	"fmt"
	"testing"
)

func TestProblems_reverseStr(t *testing.T) {
	str := "hello"
	var rStr = "olleh"
	var data []byte = []byte(str)
	s := reverseStr(data)
	if string(s) != rStr {
		t.Error("测试不通过")
	}
}

func TestProblems_twoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	target := 9
	idx := twoSum1(nums, target)
	if idx[0] != 0 || idx[1] != 1 {
		t.Error("error")
	}
}

func TestProblems_twoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 5}
	target := 9
	idx := twoSum2(nums, target)
	if idx[0] != 0 || idx[1] != 1 {
		t.Error("error")
	}
}

func TestProblems_twoSum2_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	idx := twoSum2(nums, target)
	fmt.Println(idx)
	if idx[0] != 1 || idx[1] != 2 {
		t.Error("error")
	}
}

func TestProblems_addTwoNumbers(t *testing.T) {
	fNode3 := &ListNode{Val: 3, Next: nil}
	fNode2 := &ListNode{Val: 4, Next: fNode3}
	fNode := &ListNode{Val: 2, Next: fNode2}

	sNode3 := &ListNode{Val: 4, Next: nil}
	sNode2 := &ListNode{Val: 6, Next: sNode3}
	sNode := &ListNode{Val: 5, Next: sNode2}

	resNode := addTwoNumbers(fNode, sNode)
	fmt.Println("resNode: ", resNode)
}

func TestProblems_lengthOfLongestSubstring(t *testing.T) {
	// str1 := "abcabcbb"
	// len1 := lengthOfLongestSubstring(str1)
	// fmt.Println("len1: ", len1)

	str2 := "bbbbb"
	len2 := lengthOfLongestSubstring(str2)
	fmt.Println("len2: ", len2)

	// fmt.Println("au len:", lengthOfLongestSubstring("au"))

	// str3 := "pwwkew"
	// len3 := lengthOfLongestSubstring(str3)
	// fmt.Println("len3: ", len3)
}

func TestProblems_findMaxConsecutiveOnes(t *testing.T) {
	arr1 := []int{1, 1, 0, 1, 1, 1}
	n := findMaxConsecutiveOnes(arr1)
	fmt.Println(n)
}

func TestProblems_findPoisonedDuration(t *testing.T) {
	arr1 := []int{1, 4}
	arr2 := []int{1, 2}
	arr3 := []int{1, 2, 3, 4, 5}
	d1 := 2
	d2 := 5
	fmt.Println(findPoisonedDuration(arr1, d1))
	fmt.Println(findPoisonedDuration(arr2, d1))
	fmt.Println(findPoisonedDuration(arr3, d2))
}

func TestProblems_findRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println(num)
}

func TestProblems_findRepeatNumber2(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber2(nums)
	fmt.Println(num)
}

func TestProblems_largestPerimeter(t *testing.T) {
	s := []int{1, 2, 1}
	l := largestPerimeter(s)
	fmt.Println(l)
}
