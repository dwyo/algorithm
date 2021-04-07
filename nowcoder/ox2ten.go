package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nums []string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		nums = append(nums, line)
	}
	for _, v := range nums {
		n, _ := strconv.ParseInt(v[2:], 16, 32)
		fmt.Printf("%d\n", n)
	}
}

//func Hex2Dec(n string) int64{
//	nn, _ := hex.
//	x, _ := strconv.ParseInt(string(nn), 12, 64)
//	return x
//}
