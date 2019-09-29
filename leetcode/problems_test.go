package leetcode

import (
	"fmt"
	"testing"
)

func TestProblems_reverse(t *testing.T) {
	str := "hello"
	var reverseStr = "olleh"
	var data []byte = []byte(str)
	s := reverse(data)
	if string(s) != reverseStr {
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

func TestProblems_lengthOfLongestSubstring(t *testing.T) {
	str1 := "abcabcbb"
	fmt.Println("str1: ", str1)
	str2 := "bbbbb"
	fmt.Println("str2: ", str2)
	str3 := "pwwkew"
	fmt.Println("str3: ", str3)
}
