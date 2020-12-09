package search

import (
	"fmt"
	"testing"
)

func TestBinary_bSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	idx := bSearch(nums, 14)
	if idx != 13 {
		fmt.Println("idx: ", idx)
		t.Error()
	}
}
