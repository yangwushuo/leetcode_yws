/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	/*
		解题思路:
			1先找出最短的一个字符串，因为前缀最长也不会超过最短的那个字符串
			2开始循环,最短的字符串长度为边界条件，因为前缀不会超过最大前缀长度，
			 内部循环，以两个字符串交叉式进行比对，如果出现不同说明前缀止步于此
	*/

	strsLength := len(strs)
	if strsLength == 0 {
		return ""
	}

	maxPreLength := len(strs[0])
	//找出最短的哪一个字符
	for i := 1; i < strsLength; i++ {
		strLength := len(strs[i])
		if maxPreLength > strLength {
			//如果有空串，说明前缀为空串
			if strLength == 0 {
				return ""
			}
			maxPreLength = len(strs[i])
		}
	}

	//根据最大前缀为边界条件，对字符数组中每个字符进行比对

	index := 0

over:
	for ; index < maxPreLength; index++ {
		for j := 1; j < strsLength; j++ {
			if strs[j][index] != strs[j-1][index] {
				break over
			}
		}
	}

	return strs[0][0:index]
}

// @lc code=end

