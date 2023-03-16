package chain

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	lessDummy := &ListNode{}
	betterDummy := &ListNode{}
	lessHead := lessDummy
	betterHead := betterDummy

	for head != nil {
		// 小于chechNode的链接到less节点中，其余链接到better节点中
		if head.Val >= x {
			betterDummy.Next = head
			// 继续链接下一步
			betterDummy = betterDummy.Next
		} else {
			lessDummy.Next = head
			// 继续链接下一步
			lessDummy = lessDummy.Next
		}
		head = head.Next
	}
	// 如果不将后面置空，可能会出现链表环
	betterDummy.Next = nil
	lessDummy.Next = betterHead.Next

	return lessHead.Next
}

// @lc code=end
