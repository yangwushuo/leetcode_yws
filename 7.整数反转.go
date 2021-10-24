/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {

	var num int = 0

	for x != 0 {
		num = num*10 + (x % 10)
		x = x / 10
	}

	//<<代表向左位移
	if num < -(1<<31) || num > (1<<31)-1 {
		return 0
	}

	return num
}

// @lc code=end

