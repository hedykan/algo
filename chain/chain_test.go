package chain

import (
	"fmt"
	"testing"
)

func Test22(T *testing.T) {
	str := generateParenthesis(3)
	fmt.Println(str)
}

func Test141(t *testing.T) {
	node := &ListNode{
		Val:  1,
		Next: nil,
	}
	node.Next = node
	has := hasCycle(node)
	if !has {
		t.Error("error", has)
	}

	node.Next = nil
	has = hasCycle(node)
	if has {
		t.Error("error", has)
	}
}

func Test21(t *testing.T) {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 4}

	node1 := &ListNode{Val: 1}
	node1.Next = &ListNode{Val: 3}
	node1.Next.Next = &ListNode{Val: 4}

	mergeTwoLists(node, node1)
}

func Test86(t *testing.T) {
	node := CreateChain(2, 1)

	resNode := partition(node, 2)
	PrintChain(resNode)
}

func Test23(t *testing.T) {
	chainArr := []*ListNode{CreateChain(1, 2), CreateChain(3, 4)}
	resNode := mergeKLists(chainArr)
	PrintChain(resNode)
}

func Test19(t *testing.T) {
	chainArr := CreateChain(1, 2, 4, 5, 6)
	PrintChain(chainArr)

	resNode := removeNthFromEnd(chainArr, 2)
	PrintChain(resNode)
}

func Test876(t *testing.T) {
	chainArr := CreateChain(1, 2, 4, 5, 6, 7)
	res := middleNode(chainArr)
	fmt.Println(*res)
}

func Test160(t *testing.T) {
	headA := CreateChain(1, 2, 3)
	headB := CreateChain(5, 6)
	headCommon := CreateChain(10, 11)

	// 链接链表
	headA = LinkChain(headA, headCommon)
	headB = LinkChain(headB, headCommon)

	PrintChain(headA)
	PrintChain(headB)

	// 获取公共链表头
	common := getIntersectionNode(headA, headB)

	if common != headCommon {
		t.Error("error")
	}
	PrintChain(common)
	PrintChain(headCommon)
}

func Test206(t *testing.T) {
	chain := CreateChain(1, 2, 3, 4, 5)
	PrintChain(chain)
	res := reverseList(chain)
	PrintChain(res)
}

func Test92(t *testing.T) {
	chain := CreateChain(1, 2, 3, 4, 5)
	PrintChain(chain)
	res := reverseBetween(chain, 2, 4)
	PrintChain(res)
}
