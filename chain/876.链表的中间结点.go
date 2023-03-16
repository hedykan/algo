package chain

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	p1 := head
	p2 := head

	for p1 != nil && p1.Next != nil {
		// p1以speed倍速度往前访问
		p1 = p1.Next.Next
		p2 = p2.Next
	}
	return p2
}

// @lc code=end
