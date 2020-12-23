package leetcode

import (
	"fmt"
	"testing"
)

func Test387字符串中的第一个唯一字符_firstUniqChar(t *testing.T) {
	s := "loveleetcode"
	idx := firstUniqChar(s)
	fmt.Println(idx)
}
