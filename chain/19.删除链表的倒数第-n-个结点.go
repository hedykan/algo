package chain

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 使用dummy是为了边界为空的时候不报错，在这里是为了长度为n的链条删掉倒数第n项
	dummy := &ListNode{-1, head}
	// 查找倒数第n+1位，因为用了虚拟节点dummy，所以不怕越界
	x := findFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}

// 查找链表倒数k项
func findFromEnd(head *ListNode, x int) *ListNode {
	p1 := head
	// 先寻找到k位
	for i := 0; i < x; i++ {
		p1 = p1.Next
	}
	p2 := head
	// 利用p2与p1的相对位置为k，p1跑到尾部时p2为n-k+1
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

// @lc code=end
