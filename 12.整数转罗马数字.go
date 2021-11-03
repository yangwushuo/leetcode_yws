/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {

	/*
		解题思路:
			先把每个值得所对应的字符做好
			接下来循环拼接
			举例140
			140大于100小于400，所以有C，40不大于100所以需要下标减1
			40等于40 所以有XL
			举例3
			3大于1 所以有I，得2
			2大于1 所以有I，得1
			1等于1 所以有I, 得0结束
	*/

	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var buffer bytes.Buffer
	for x := len(intSlice) - 1; x >= 0 && num != 0; {
		if num >= intSlice[x] {
			buffer.WriteString(roman[x])
			num -= intSlice[x]
		} else {
			x--
		}
	}

	return buffer.String()
}

// @lc code=end

