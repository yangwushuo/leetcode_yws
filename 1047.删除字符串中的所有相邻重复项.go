/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(s string) string {

	/*
		解题思路:
			1创建一个栈用于存放字符，循环压入栈内
			2如果当前字符与栈顶一致，则删除栈顶字符，进行下一轮判断
	*/

	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)

}

// @lc code=end

