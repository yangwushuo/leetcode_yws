/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {

	/*
		解题思路:
			思路可以参照15.三叔之和也是使用双指针
			1循环体内部，判断当第一个数值与上一次循环的数值不重复，如果重复说明当前数值已经经过计算，不要再重复计算
			2如果三个数相加等于target直接返回target即可 ，因为没有比这个数更小的了
			3如果三数之和大于target因为缩小范围，这里a<b<c ,a+b+c>target,a不能动b只会增大，所以缩小c，移动右指针缩小范围
			4同理如果小于target应该增大数值，所以向右移动left
			5最后一次执行循环克定是i=-3，left=-2，right=-1，就是最后三个数，执行完也就结束整个循环体，返回最优解即可
	*/

	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n-2; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 left 和 right
		left, right := i+1, n-1
		for left < right {
			num1, num2, num3 := nums[i], nums[left], nums[right]
			sum := num1 + num2 + num3
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			//
			if sum > target {
				for left < right && num3 == nums[right] {
					right--
				}
			} else {
				for left < right && num2 == nums[left] {
					left++
				}
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// @lc code=end

