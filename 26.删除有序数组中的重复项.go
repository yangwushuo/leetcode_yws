/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	/*
		解题思路:
			1创建两个快慢指针，分别指向下标为1
			2从1位置开始遍历，判断当前快指针指向的数值是否与前一个相同，
			 如果相同不做操作，如果不相同则将快指针指向的数值赋值给慢指针，然后慢指针向前移动一位
			 注意:慢指针前面都是不重复的数值，所以结尾返回慢指针的下标即可
			 1 1 2 3 4 4 4 4 5
			   |
		    fast|slow
	*/

	if len(nums) < 2 {
		return len(nums)
	}
	//创建两个快慢双指针
	var (
		fast    = 1
		slow    = 1
		numsLen = len(nums)
	)

	for fast = 1; fast < numsLen; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// @lc code=end

