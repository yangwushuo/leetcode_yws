/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {

	/*
		解题思路
			1此题解题思路和三数之和一样先排序
			2只是此题是四数之和
			3遍历数组，确定第一个数字，然后再遍历第一个数字下标以后得数组，确定第二个数(一般是第一个数字下标后一个)。
			 再确定左右指针，不断缩小范围进行比对
	*/

	sort.Ints(nums)
	var (
		result     = [][]int{}
		numsLength = len(nums)
	)

	for i := 0; i < numsLength-3; i++ {
		first := nums[i]
		if i != 0 && first == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsLength-2; j++ {
			second := nums[j]
			if j != i+1 && second == nums[j-1] {
				continue
			}
			left := j + 1
			right := numsLength - 1
			for left != right {
				num1, num2 := nums[left], nums[right]
				if first+second+num1+num2 == target {
					result = append(result, []int{first, second, nums[left], nums[right]})
					for left < right && num1 == nums[left] {
						left++
					}
					for left < right && num2 == nums[right] {
						right--
					}
				} else if first+second+nums[left]+nums[right] > target {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}

// @lc code=end

