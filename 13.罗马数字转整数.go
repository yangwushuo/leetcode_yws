/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {

	/*
		解题思路:
			1建立map 罗马数字 -> 数字
			2遍历字符串，如果当前字符映射的数字小于后一位字符映射的数字，则用总和减去当前字符映射的数字。
						如果不小于说明不是特殊数字，则直接加入总和就可以
	*/

	var res int
	symbol := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := 0; i < len(s); i++ {
		val := symbol[s[i]]
		if i < len(s)-1 && symbol[s[i+1]] > val {
			res -= val
		} else {
			res += val
		}
	}
	return res
	//
}

// @lc code=end

