package chain

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode

// 反转开头后n个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	// 获取第n个节点作为头返回
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	// 将尾部不断设置到当前的节点尾部，直到递归结束将会设置到原来的头节点的尾部
	head.Next = successor
	return last
}

// @lc code=end
