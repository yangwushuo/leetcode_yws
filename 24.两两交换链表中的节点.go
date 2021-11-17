/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	/*
		解题思路:(迭代)
			也可以通过迭代的方式实现两两交换链表中的节点。
			创建哑结点 dummyHead，令 dummyHead.next = head。令 temp 表示当前到达的节点，初始时 temp = dummyHead。
			每次需要交换 temp 后面的两个节点。
			如果 temp 的后面没有节点或者只有一个节点，则没有更多的节点需要交换，因此结束交换。
			否则，获得 temp 后面的两个节点 node1 和 node2，通过更新节点的指针关系实现两两交换节点。
			具体而言，交换之前的节点关系是 temp -> node1 -> node2，交换之后的节点关系要变成 temp -> node2 -> node1，因此需要进行如下操作。
			temp.next = node2
			node1.next = node2.next
			node2.next = node1
			完成上述操作之后，节点关系即变成 temp -> node2 -> node1。再令 temp = node1，对链表中的其余节点进行两两交换，直到全部节点都被两两交换。
			两两交换链表中的节点之后，新的链表的头节点是 dummyHead.next，返回新的链表的头节点即可。
	*/

	var (
		dumpHead = &ListNode{
			Next: head,
		}
		temp = dumpHead
	)

	for temp.Next != nil && temp.Next.Next != nil {
		var (
			node1 = temp.Next
			node2 = temp.Next.Next
		)
		//改变次序
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}

	return dumpHead.Next
}

// @lc code=end

