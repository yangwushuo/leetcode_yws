/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	/*
		解题思路:
			I.给定有序二维数组，只需要确定目标数所在行内范围，再从目标行中找出目标数
			II.	以二维数组左下角为原点，建立直角坐标轴。
				若当前数字大于了查找数，查找往上移一位。
				若当前数字小于了查找数，查找往右移一位。
			III.将二维数组(有序)转为一维数组(升序)，再用二分查找法找出目标数
	*/

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	c := len(matrix[0]) - 1
	for r := 0; r < len(matrix); r++ {
		if matrix[r][c] >= target {
			for i := 0; i <= c; i++ {
				if matrix[r][i] == target {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=end

