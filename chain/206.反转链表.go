package chain

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
	if head == nil || head.Next == nil {
		return head
	}
	// 最尾节点被返回
	last := reverseList(head.Next)
	// 自身链接到下一个节点后面
	head.Next.Next = head
	// 将尾部不断设置到当前的节点尾部，直到递归结束将会设置到原来的头节点的尾部
	head.Next = nil
	return last
}

// @lc code=end
