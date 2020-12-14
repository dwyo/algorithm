package leetcode

import (
	"fmt"
	"testing"
)

func Test49字母异位词分组_groupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// strs := []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	s := groupAnagrams(strs)
	fmt.Println(s)
}
