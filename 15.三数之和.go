/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {

	/*
		代码思路参照:https://leetcode-cn.com/problems/3sum/solution/dai-ma-sui-xiang-lu-dai-ni-gao-ding-ha-x-w6pp/
		解题思路:双指针法
			1先排序，如果长度小于3直接放回数组
			2开始循环数组，当排序数组中有大于0说明后面的数也大于0不用再继续循环
			3判断当前数字是否与前一个重复，如果与前一个重复说明当前数字数值已经与后面数值进行计算过，防止重复计算被加入结果数组中
			4确定左右指针分别指向的数值f
			 (1)如果3数之和等于0直接创建数组压入结果数组中
			 然后再考虑移动左指针和右指针，这里移动指针要考虑到移动完成后数值是否还是和原来一样
			 如果一样不能再去计算，否则会造成数组中有重复结果，应该再继续移动指针直到与原来数值不等
			 (2)如果3数之和大于0，右指针需要向前移动缩小数值之和
			 (3)同理如果小于0，左指针需要向后移动增大数值之和
			5直到左右指针相等，说明第二个数和第三个数已经寻找完毕，需要前移第一个数。以此循环知道遍历到倒数第三个数，
			 此时左指针是num2前一个是num1，后一个是num3
	*/

	//判断长度
	numsLen := len(nums)
	sumToZero := [][]int{}
	sort.Ints(nums)

	if numsLen < 3 {
		return sumToZero
	}

	//减2，当循环到倒数第三个的时候，后面还有两个指针所以减去2
	for i := 0; i < numsLen-2; i++ {
		num1 := nums[i]
		//如果大于0那么后面的数相加都大于0
		if num1 > 0 {
			break
		}
		//重复数不列入计算
		if i > 0 && num1 == nums[i-1] {
			continue
		}
		//定义左右指针
		left, right := i+1, numsLen-1
		for left < right {
			num2, num3 := nums[left], nums[right]
			if num1+num2+num3 == 0 {
				sumToZero = append(sumToZero, []int{num1, num2, num3})
				//移动左右指针
				for left < right && num2 == nums[left] {
					left++
				}
				for left < right && num3 == nums[right] {
					right--
				}
			} else if num1+num2+num3 > 0 {
				//如果三个数(从小到大排序)相加大于0说明，该移动右指针
				right--
			} else {
				//如果三个数(从小到大排序)相加小于0说明，该移动左指针
				left++
			}
		}
	}
	return sumToZero
}

// @lc code=end

