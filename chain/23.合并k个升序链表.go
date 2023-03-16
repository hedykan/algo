package chain

import "container/heap"

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
	if len(lists) == 0 {
		return nil
	}
	dummy := &ListNode{}
	p := dummy

	pq := make(PriorityQueue, 0)
	heap.Init(&pq) // 使用最小堆

	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head) // 把三个头压入堆中，等待取最小
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next) // 把新的节点压入堆中，等待取最小
		}
		p = p.Next
	}

	return dummy.Next
}

// 需要实现sort接口和push,pop接口
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// @lc code=end
