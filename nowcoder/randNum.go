package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main()  {
	var nums [][]int
	var num []int
	input := bufio.NewScanner(os.Stdin)
	first := 0
	cnt := 0
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		n, _ := strconv.Atoi(line)
		if first == 0 {
			first = n
			continue
		}
		cnt++
		num = append(num, n)
		if cnt == first {
			nums = append(nums, num)
			num = []int{}
			cnt = 0
			first = 0
		}
	}
	for _,num := range nums {
		randNum(num)
	}
}

func randNum(nums []int){
	temp := make([]int, 1001)
	for _, num := range nums {
		temp[num]++
	}
	for k,v :=range temp {
		if v != 0 {
			fmt.Println(k)
		}
	}
}
