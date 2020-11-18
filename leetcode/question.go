package leetcode

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	tmp := 0
	for j := 0; j < len(nums); j++ {

		if tmp == sum-nums[j]-tmp {
			return j
		}
		tmp += nums[j]
	}

	return -1
}

func pivotIndex2(nums []int) int {
	sum := 0
	for i, _ := range nums {
		sum += nums[i]
	}
	tmp := 0
	cnt := len(nums)
	for j := 0; j < cnt; j++ {

		if tmp*2+nums[j] == sum {
			return j
		}
		tmp += nums[j]
	}

	return -1
}
