package leetcode

import (
	"fmt"
	"testing"
)

func Test435无重叠区间_eraseOverlapIntervals(t *testing.T) {
	nums := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(nums))
}
