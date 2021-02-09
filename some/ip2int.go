package some

import (
	"strconv"
	"strings"
)

func Ip2int(ip string) int {
	ipStr := strings.Split(ip, ".")
	var ret int
	n := len(ipStr)
	for i := n - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(ipStr[i])
		if i == 3 {
			ret = num
		} else {
			ret |= num << (8 * (n - i - 1))
		}
	}
	return ret
}

func Ip2int1(ip string) int {
	ipStr := strings.Split(ip, ".")
	ip3, _ := strconv.Atoi(ipStr[3])
	ip2, _ := strconv.Atoi(ipStr[2])
	ip1, _ := strconv.Atoi(ipStr[1])
	ip0, _ := strconv.Atoi(ipStr[0])

	ret := ip3 | ip2<<8 | ip1<<16 | ip0<<24

	return ret
}
