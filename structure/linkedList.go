package structure

import (
	"errors"
	"fmt"
)

type nodeValType = int

type LinkedList struct {
	Val nodeValType
	Next *LinkedList
}
// CreateLinkedList 创建一个链表
func CreateLinkedList() LinkedList {
	return LinkedList{0, nil}
}
// Len 链表长度
func (head *LinkedList) Len() int {
	length := 0
	p := head.Next
	for p != nil {
		length++
		p = p.Next
	}
	return length
}
// PushNode 链表结尾加入结点, 不可成为环
func (head *LinkedList) PushNode(new ...*LinkedList) {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	for _, v := range new {
		p.Next = v
		p = p.Next
	}
	p.Next = nil // 防止最后一个结点指向了前面的某个结点形成了环
	return
}
// PushVal 和PushNode一样，但参数为val
func (head *LinkedList) PushVal(new ...nodeValType) {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	for _, v := range new {
		p.Next = &LinkedList{v, nil}
		p = p.Next
	}
	return
}
// Delete 删除结点
func (head *LinkedList) Delete(val nodeValType) error {
	if head.Next == nil {
		return errors.New("Cann't found the node")
	}
	pDeleted := head.Next
	pDeletedPre := head
	for pDeleted != nil {
		if pDeleted.Val == val {
			pDeletedPre.Next = pDeleted.Next
			return nil
		}
		pDeleted = pDeleted.Next
		pDeletedPre = pDeletedPre.Next
	}
	return errors.New("Cann't found the node")
}
// Put 修改结点值
func (head *LinkedList) Put(val, newVal nodeValType) error {
	p := head.Next
	for p != nil {
		if p.Val == val {
			p.Val = newVal
			return nil
		}
		p = p.Next
	}
	return errors.New("Cann't found the node")
}
// PrintList 输出链表
func (head *LinkedList) PrintList() {
	p := head.Next
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
	return
}
// hasCycle 检查是否有环
func (head *LinkedList) HasCycle() bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head.Next
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}
// MiddleNode 返回中间结点，如果有两个中间结点则返回第二个
func (head *LinkedList) MiddleNode() *LinkedList {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
// mergeTwoLists 合并两个有序的链表
func MergeTwoLists(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *LinkedList
	if l1.Val >= l2.Val{
		res = l2
		res.Next = MergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = MergeTwoLists(l1.Next,l2)
	}
	return res
}
// removeNthFromEnd 删除链表的倒数第n个结点
func RemoveNthFromEnd(head *LinkedList, n int) *LinkedList {
	dummy := &LinkedList{Next: head}
	first := dummy
	second := dummy
	for n >= 0 {
		first = first.Next
		n--
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
// ReverseList 反转链表
func ReverseList(head *LinkedList) *LinkedList {
	var pre *LinkedList
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
// 移除链表重复元素(链表无序), 使用哈希记录或者双重循环