package nowcoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
链接：https://www.nowcoder.com/questionTerminal/fe298c55694f4ed39e256170ff2c205f
来源：牛客网

有这样一道智力题：“某商店规定：三个空汽水瓶可以换一瓶汽水。
小张手上有十个空汽水瓶，她最多可以换多少瓶汽水喝？”答案是5瓶，
方法如下：
	先用9个空瓶子换3瓶汽水，喝掉3瓶满的，喝完以后4个空瓶子，
	用3个再换一瓶，喝掉这瓶满的，这时候剩2个空瓶子。
	然后你让老板先借给你一瓶汽水，喝掉这瓶满的，
	喝完以后用3个空瓶子换一瓶满的还给老板。如果小张手上有n个空汽水瓶，最多可以换多少瓶汽水喝？
 */

func main()  {
	var nums []int
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		num, _ := strconv.Atoi(line)
		if num == 0 {
			break
		}
		nums = append(nums, num)
	}
	for _,v := range nums {
		fmt.Println(soda(v))
	}
}

func soda(n int) int{
	cnt := 0
	flag := 0
	for n > 1 {
		if n == 2 {
			flag++
			n += 1
		}
		a := n %3
		b := n / 3
		n = a + b
		cnt += b
	}
	return cnt
}
