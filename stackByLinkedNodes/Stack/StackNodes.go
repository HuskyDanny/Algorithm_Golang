package StackByNodes

import (
	"fmt"
)

type node struct {
	next  *node
	val interface{}
}

type StackByNodes struct {
	topNode *node
}


func StackInit() *StackByNodes{
	return &StackByNodes{nil}
}

func (this *StackByNodes) Push(val interface{}) {
	this.topNode = &node{next : this.topNode, val : val}
}

func (this *StackByNodes) Pop() interface{} {
	value := this.topNode.val
	this.topNode = this.topNode.next
	return value
}

func (this *StackByNodes) PrintAll() {
	
	temp := this.topNode

	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
}
