package chain

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	// 如果p1==p2
	for p1 != p2 {
		// 两个指针都走一遍两个链表，他们所经过路径相同，即A'+Common+B' == B'+Common+A'
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2先走一段B在走一段A
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// @lc code=end
