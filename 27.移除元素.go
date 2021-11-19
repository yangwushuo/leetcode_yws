/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:len(nums)]...)
			i--
		}
	}
	return len(nums)
}

// @lc code=end

