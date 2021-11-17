/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	/*
			解题思路:
				1 dumpHead头节点前一个哑节点，temp始终是指向k周期链表的节点
				  n1为k周期起始节点，n2为k周期终节点
				2 循环移动n2 如果为空代表结束 内部判断为判断是否达到周期k
					如果n2移动达到周期k则开始颠倒周期内的链表 周期内的链表，头变尾  尾变头 所以返回新的头也就是尾


				nil ->  1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> nil
				|	    |
		dumpHead|temp n1|n2
	*/

	var (
		count    = 1
		dumpHead = &ListNode{
			Next: head,
		}
		temp = dumpHead
		n1   = head
		n2   = head
	)

	for n2 != nil {
		if count == k {
			//记住下一个周期的起始节点
			nextN1 := n2.Next

			//相面两行起到将颠倒后的链表重新添加到总链表中
			//颠倒当前周期内的链表获取新得头部，将上一个周期链表指向当前新的链表头部
			temp.Next = reversal(n1, n2)
			//将当前周期与下一个周期连接起来
			n1.Next = nextN1

			//更新temp n1 n2 count开始下一个周期
			temp = n1
			n1 = nextN1
			n2 = nextN1
			count = 1
		} else {
			count++
			n2 = n2.Next
		}
	}

	return dumpHead.Next
}

//颠倒链表返回新的头节点
func reversal(node1 *ListNode, node2 *ListNode) *ListNode {
	var front *ListNode
	var back *ListNode

	//node1，node2分别是头尾节点
	for node1 != node2 {
		//记住后一个节点
		back = node1.Next
		//将当前next指向改成前一个节点
		node1.Next = front

		//front和node1分别向前移动
		front = node1
		node1 = back
	}
	node1.Next = front

	return node2
}

// @lc code=end

