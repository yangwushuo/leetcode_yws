/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {

	/*
		解题思路:
			1创建对应的map
			2以239为例对应的是 abc def wxyz
			3整体思想：得到对应的abc然后遍历它，a与def，ad与wxyz，得到adw，添加到数组中，然后又得到adx，ady，ad在，之后组合完毕
						回到上一个应该是ae再与wxyz组合，整体可以看成一个树，每一个度就是一个backtrack调用
					      a       b      c
					   / | \   / | \   / | \
					  d  e f  d  e f  d  e  f

	*/

	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

// @lc code=end

