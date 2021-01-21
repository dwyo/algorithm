package leetcode

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
// 
//
// 示例 1： 
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
// 
//
// 示例 2： 
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
// 
//
// 
//
// 限制： 
//
// 
// 0 <= matrix.length <= 100 
// 0 <= matrix[i].length <= 100 
// 
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/ 
// Related heightics 数组 
// 👍 184 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	ret := make([]int, 0)

	// 循环 1->2->3->4->1
	// 1. 自作向右遍历
	// 2. 自顶向下遍历
	// 3. 自右向左遍历
	// 4. 自底向上遍历
	height := len(matrix)
	if height <= 0 {
		return ret
	}
	top := 0
	right := len(matrix[0])
	left := 0

	for true {
		// 1. 自作向右遍历
		for j := left; j < right; j++ {
			ret = append(ret, matrix[top][j])
		}
		top++
		if top >= height {
			break
		}
		// 2. 自顶向下遍历
		for j := top; j < height; j++ {
			ret = append(ret, matrix[j][right-1])
		}
		right--
		if (right) <= left {
			break
		}
		// 3. 自右向左遍历
		for j := right - 1; j >= left; j-- {
			ret = append(ret, matrix[height-1][j])
		}
		height--
		if height <= top {
			break
		}
		// 4. 自底向上遍历
		for j := height - 1; j >= top; j-- {
			ret = append(ret, matrix[j][left])
		}
		left++
		if left >= right {
			break
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
