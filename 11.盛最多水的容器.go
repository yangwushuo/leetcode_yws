/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {

	/*
		解题思路:
			使用双指针
			1左指针指向头部，右指针指向尾部
			2循环，判断条件左右指针不能相同
			3判断面积
			4移动指针，只移动最短的那一根的指针
				(原理如果移动长度最长的那一根指针，则面积只会减小而不会增大，如果移动最短的那一根面积则会增大或减小)
				面积: 底(每次缩小1)*高(移动最短指针之后，面积有两种可能增大或减小，但如果移动最长的指针则面积只会减小)
	*/

	maxA := 0
	l := 0
	r := len(height) - 1

	for l != r {
		area := Min(height[l], height[r]) * (r - l)
		if area > maxA {
			maxA = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxA
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

