/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {

	/*
		解题思路:动态规划
		1判断字符串是否小于2如果小于则直接返回
		2创建二维数组用于存放字符串i到j之间的是否是回文数 默认单字符为true
		3两个循环，第一个循环控制长度，第二个循环控制循环次数

	*/
	//1
	maxLen := 1
	begin := 0
	length := len(s)
	if length < 2 {
		return s
	}
	//2
	result := make([][]bool, length)
	for i := 0; i < length; i++ {
		result[i] = make([]bool, length)
		result[i][i] = true
	}

	//3
	for i := 2; i <= length; i++ {
		for j := 0; j < length; j++ {
			//算出步长
			k := j + i - 1
			if k > length-1 {
				break
			}

			//判断两个字符是否相等
			if s[j] == s[k] {
				//判断是否是长度2或者3的回文数 如果不是则判断两个字符前一个和后一个是否相等，如果相等说明是回文数
				if k-j < 3 {
					result[j][k] = true
				} else if result[j+1][k-1] == true {
					result[j][k] = true
				}

				//判断长度是否是最长，且是回文数
				if k-j+1 > maxLen && result[j][k] == true {
					maxLen = k - j + 1
					begin = j
				}
			}
		}
	}

	return s[begin : begin+maxLen]
}

// @lc code=end

/*
	最佳解答:
	func longestPalindrome(s string) string {
		if s == "" {
			return ""
		}
		start, end := 0, 0
		for i := 0; i < len(s); i++ {
			left1, right1 := expandAroundCenter(s, i, i)
			left2, right2 := expandAroundCenter(s, i, i + 1)
			if right1 - left1 > end - start {
				start, end = left1, right1
			}
			if right2 - left2 > end - start {
				start, end = left2, right2
			}
		}
		return s[start:end+1]
	}

	func expandAroundCenter(s string, left, right int) (int, int) {
		for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
		return left + 1, right - 1
	}
*/