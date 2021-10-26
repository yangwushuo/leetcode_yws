/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {

	/*
		解题思路:
			1看作V字循环，算出循环周期为n*numRows-2个周期
			2遍历字符
	*/

	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	n := 2*numRows - 2
	for i, c := range s {
		x := i % n
		rows[min(x, n-x)] += string(c)
	}
	return strings.Join(rows, "")
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

