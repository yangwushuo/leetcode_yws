/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	/*
		解题思路:
			遍历haystack每个字符，如果当前字符等于needle首字母，
			则进一步比对haystack当前下标加上needle的长度的字符串，与needle整个字符串进行比对，如果相同则放回下标

	*/

	var (
		haystackLen = len(haystack)
		needleLen   = len(needle)
	)

	if needleLen == 0 {
		return 0
	} else if haystackLen < needleLen {
		return -1
	}

	for i := 0; i < haystackLen; i++ {
		if haystack[i] == needle[0] && haystackLen-i >= needleLen {
			if haystack[i:i+needleLen] == needle {
				return i
			}
		}
	}

	return -1
}

// @lc code=end

