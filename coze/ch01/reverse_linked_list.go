package ch01

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Run_reverseList() {
	head := buildLinkedList()
	printLinkedList(head)
	reversed := reverseList(head)
	printLinkedList(reversed)
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode

	cur = head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev

		prev = cur
		cur = tmp
	}
	return prev
}

func buildLinkedList() *ListNode {
	var one = &ListNode{Val: 1}
	var two = &ListNode{Val: 2}
	var three = &ListNode{Val: 3}
	var four = &ListNode{Val: 4}

	var head = one
	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = nil

	return head
}

func printLinkedList(head *ListNode) {
	cur := head

	fmt.Println("printLinkedList start")
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
	fmt.Println("printLinkedList end")

}
