package chain

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// var nodeMap map[*ListNode]bool

func hasCycle(head *ListNode) bool {
	return hasCycleItem1(head)
	// 到头了
	// nodeMap = make(map[*ListNode]bool)
	// return hasCycleItem(head)
}

// func hasCycleItem(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	// 有环，命中缓存了
// 	if _, ok := nodeMap[head]; ok {
// 		return true
// 	} else {
// 		nodeMap[head] = true
// 		return hasCycleItem(head.Next)
// 	}
// }

func hasCycleItem1(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=end
