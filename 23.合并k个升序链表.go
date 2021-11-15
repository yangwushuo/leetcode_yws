/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	/*
		解题思路:
			k个链表的合并，可以看做是k-1次，每两个链表之间的合并
			创建一个头尾都是空的节点，如果两个链表循环完毕则结束merage函数放回头节点的下一个节点(真正的头节点)
			之后再获取下一个链表的头节点和前面排好序的头节点进行排序。最终遍历完数组获得一个k个链表合并之后的头节点
			注意:每两个链表进行合并
	*/

	var pre, cur *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = merge(pre, cur)
	}
	return pre
}

func merge(firstList *ListNode, secondList *ListNode) *ListNode {

	tail := &ListNode{}
	head := tail
	//
	for firstList != nil || secondList != nil {
		if firstList != nil && secondList != nil {
			if firstList.Val < secondList.Val {
				tail.Next = firstList
				tail = firstList
				firstList = firstList.Next
				for (firstList != nil) && (firstList.Val < secondList.Val) {
					tail = firstList
					firstList = firstList.Next
				}
			} else {
				tail.Next = secondList
				tail = secondList
				secondList = secondList.Next
				for (secondList != nil) && (secondList.Val < firstList.Val) {
					tail = secondList
					secondList = secondList.Next
				}
			}
		} else if firstList == nil {
			tail.Next = secondList
			break
		} else if secondList == nil {
			tail.Next = firstList
			break
		}
	}

	return head.Next
}

// @lc code=end

