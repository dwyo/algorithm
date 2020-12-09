package leetcode

import (
	"fmt"
	"testing"
)

func TestQuestion_pivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	idx1 := pivotIndex(nums)
	if idx1 != 3 {
		fmt.Println("idx1: ", idx1)
		t.Fail()
	}

	nums1 := []int{1, 2, 3}
	idx2 := pivotIndex(nums1)
	if idx2 != -1 {
		fmt.Println("idx2: ", idx2)
		t.Fail()
	}
}

func TestQuestion_pivotIndex2(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	idx1 := pivotIndex2(nums)
	if idx1 != 3 {
		fmt.Println("idx1: ", idx1)
		t.Fail()
	}

	nums1 := []int{1, 2, 3}
	idx2 := pivotIndex2(nums1)
	if idx2 != -1 {
		fmt.Println("idx2: ", idx2)
		t.Fail()
	}
}

func TestQuestion_canCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	idx := canCompleteCircuit(gas, cost)
	if idx != 3 {
		fmt.Println("idx: ", idx)
		t.Fail()
	}
}

func TestQuestion_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	for _, a := range nums {
		fmt.Println("value: ", a)
	}
}

func TestQuestion_countPrimes(t *testing.T) {
	cnt1 := countPrimes(1500000)
	fmt.Println("cnt1: ", cnt1)
}

func TestQuestion_matrixScore(t *testing.T) {
	// A := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
	// A := [][]int{{0}, {1}}
	// A := [][]int{{0, 1}, {1, 1}}
	A := [][]int{{0, 1, 1}, {1, 1, 1}, {0, 1, 0}}
	sum := matrixScore(A)
	fmt.Println(sum)
}
