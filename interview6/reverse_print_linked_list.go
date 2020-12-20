package interview6

import "container/list"

/*

题目：从尾到头打印链表
输入一个链表，按链表从尾到头的顺序返回一个ArrayList。
示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*

解法1：
从尾到头返回链表的结点，可以想到在遍历的过程中，将后遍历到的元素先放入数组中，那么就能实现题目要求的效果，可以想到使用递归来完成。不过递归的效率并不好。

*/

func solution1(head *ListNode) []int {
	if head == nil {
		return nil
	}

	return appendData(head)
}

func appendData(node *ListNode) []int {
	if node.Next != nil {
		list := appendData(node.Next)
		list = append(list, node.Val)
		return list
	}

	return []int{node.Val}
}

/*
解法2：
考虑到递归的效率很差，可以考虑先将链表反转，再依次将反转后的元素放入slice返回即可。
*/
func solution2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	reversed := reverseLinkedList(head)
	var res []int

	for reversed != nil {
		res = append(res, reversed.Val)
	}

	return res
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}

	return newHead
}

/*
解法3：
解法2先反转链表再将元素依次放入slice返回，因此也可以考虑直接将链表元素依次放入slice，在slice中进行反转，这样仅仅需要遍历一半的slice元素即可。
*/

func solution3(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var res []int
	for head != nil {
		res = append(res, head.Val)
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i--
		j++
	}

	return res
}

/*
解法4：
考虑到本题可以使用递归完成，那么结合go本身的defer函数的特性，可以考虑利用该特性完成解题。
*/
func solution4(head *ListNode) (res []int) {
	for head != nil {
		tmp := head
		defer func() {
			res = append(res, tmp.Val)
		}()
		head = head.Next
	}

	return
}

/*
解法5：
递归本身就是一种栈结构，因此考虑使用栈这种数据结构来完成。
*/
func solution5(head *ListNode) []int {
	if head == nil {
		return nil
	}

	stack := list.New()
	for head != nil {
		stack.PushFront(head.Val)
	}

	var res []int
	for elm := stack.Front(); elm != nil; elm = elm.Next() {
		res = append(res, elm.Value.(int))
	}

	return res
}
