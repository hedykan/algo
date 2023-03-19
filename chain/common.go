package chain

import (
	"errors"
	"fmt"
)

// 通过数组的形式打印链表
func PrintChain(head *ListNode) error {
	var res []int
	var i = 0
	for i = 0; i < 256; i++ {
		if head == nil {
			break
		}
		res = append(res, head.Val)
		head = head.Next
	}
	if i == 256 {
		return errors.New("too long")
	}
	fmt.Println(res)
	return nil
}

// 创建链表
func CreateChain(intArr ...int) *ListNode {
	dummy := &ListNode{}
	head := dummy
	for i := 0; i < len(intArr); i++ {
		dummy.Next = &ListNode{Val: intArr[i]}
		dummy = dummy.Next
	}
	dummy.Next = nil
	return head.Next
}

// 链接两个链表
func LinkChain(headA, headB *ListNode) *ListNode {
	p1 := headA
	for p1.Next != nil {
		p1 = p1.Next
	}
	p1.Next = headB
	return headA
}

// 查找链表倒数k项,这里实际上使用的是加法
func FindFromEnd(head *ListNode, x int) *ListNode {
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

// 可以利用快慢指针做乘除法，如查找第n/2个节点
func FindDivNNode(head *ListNode, speed int) *ListNode {
	p1 := head
	p2 := head

	for p1.Next != nil {
		// p1以speed倍速度往前访问
		for i := 0; i < speed; i++ {
			if p1 == nil {
				break
			}
			p1 = p1.Next
		}
		p2 = p2.Next
	}
	return p2
}
