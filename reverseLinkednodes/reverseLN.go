package reverseLinkednodes

import "fmt"

type LinkedNode struct {
	next *LinkedNode
	val int
}

func (p *LinkedNode) PrintAll() {
	dummy := &LinkedNode{nil, 0}
	dummy.next = p

	for dummy.next != nil {
		dummy = dummy.next
		fmt.Println(dummy.val)
	}
}

func ReverseIterative(head *LinkedNode) *LinkedNode{

	var pre *LinkedNode = nil 
	

	for head != nil {
		
		nextHold := head.next
		
		head.next = pre
		pre = head
		head = nextHold
		
	}

	return pre
	
}



func ReverseRecursive(head *LinkedNode) *LinkedNode {

	var res *LinkedNode
	var innerReverse func(head *LinkedNode) *LinkedNode
	innerReverse = func(head *LinkedNode) *LinkedNode {

		if head.next == nil {
			res = head
			return head
		}
		
		backNode := innerReverse(head.next)
		backNode.next = head
		
		return head
	}

	innerReverse(head)
	head.next = nil

	return res
}