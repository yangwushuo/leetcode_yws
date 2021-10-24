/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	/*
		解题思路:
			1创建两个指针pre指向nil头部前一个虚无指针
			cur一个头指针，指向头部。
			2循环判断cur头指针是否为空，如果为空则放回nil
				1创建next存放头指针下一个节点
				2改变cur.next节点方向，为前一个节点而不是后一个节点
				3向前一步移动两个指针指向
			3最后pre指针一定指向最后一个节点
	*/

	var pre *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// @lc code=end

