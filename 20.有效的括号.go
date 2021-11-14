/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {

	/*
		解题思路:
			1利用栈，如果是{[(这样得起始字符则直接压入栈内
			2如果是)]}这样的结尾字符，则需要进行判断，判断前一个字符是否是对应的起始字符，
			 如果对应则从栈末尾删除对应起始字符(对应字符永远在栈末尾)，然后进行下一次循环
			3最后结尾需要利用if判断栈内是否为0，如果不为0说明栈内存在起始字符，没有对应的结尾字符，则也是错
	*/

	var (
		result = true
		strLen = len(s)
		stack  = []uint8{}
	)

	if strLen < 2 {
		return false
	}

	for i := 0; i < strLen; i++ {
		if s[i] == 123 || s[i] == 91 || s[i] == 40 {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) != 0 && ((s[i] == 125 && stack[len(stack)-1] == 123) ||
			(s[i] == 93 && stack[len(stack)-1] == 91) ||
			(s[i] == 41 && stack[len(stack)-1] == 40)) {
			stack = stack[:len(stack)-1]
		} else {
			result = false
			break
		}
	}

	if len(stack) != 0 {
		result = false
	}

	return result
}

// @lc code=end

