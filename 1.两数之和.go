/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	/*
		解题思路：
		1创建一个map类型用于存放 值对应下标
		2循环数值数组，用target-num作为键去判断map是否存在值(对应下标)
		3如果不存在则将num作为键index作为值存储到map中
		4如果存在target-num这个键值则数组中含有两个值相加等于target
		5将两个数对应下标添加到result中并返回
		此题关键再于map存储值对索引
	*/

	maps := make(map[int]int)
	result := make([]int, 0)

	for index, num := range nums {
		_, isTrue := maps[target-num]
		if isTrue {
			result = append(result, maps[target-num])
			result = append(result, index)
			break
		}
		maps[num] = index
	}
	return result
}

// @lc code=end

