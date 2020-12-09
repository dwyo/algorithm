package search

func bSearch(nums []int, val int) int {
	length := len(nums)
	left := 0
	right := length - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
