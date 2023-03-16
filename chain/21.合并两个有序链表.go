package chain

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	// l1头节点
	p1 := list1
	// l2头节点
	p2 := list2

	// p用来遍历指针, dummy保存开始的指针
	p := dummy

	for p1 != nil && p2 != nil {
		// p1的值大于
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// 指针不断前进，进入刚连接上的节点
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}

// @lc code=end
