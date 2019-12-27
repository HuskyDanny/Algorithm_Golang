package MyHeap

import (
	"fmt"
	
)

//Index starts at 1
type MyHeap struct {
	arr []int
	length int
	
}

func NewMyHeap(capacity int) *MyHeap {
	if capacity == 0 {
		return nil
	}
	return &MyHeap{make([]int, capacity), 0}	
}

func (this *MyHeap) Append(val int) bool {

	this.length++
	this.arr[this.length] = val
	targetIndex := this.percolateUp(val)
	this.swap(this.length, targetIndex)

	return true
}

func (this *MyHeap) RemoveTop() int {
	topValue := this.arr[1]
	this.swap(1, this.length)
	this.length -= 1

	targetIndex := this.percolateDown(topValue)

	this.swap(1, targetIndex)

	return topValue
}

func (this *MyHeap) percolateUp(val int) int {

	startIndex := (this.length )

	for startIndex > 1  && this.arr[startIndex / 2] < val {
		startIndex = startIndex / 2
	}

	return startIndex
	
}

func (this *MyHeap) swap(indexI int, indexJ int) {
	temp := this.arr[indexI]
	this.arr[indexI] = this.arr[indexJ]
	this.arr[indexJ] = temp
}

func (this *MyHeap) percolateDown(val int) int {
	
	startIndex := 1

	for startIndex <= this.length && (val < this.arr[startIndex * 2] || val < this.arr[startIndex * 2 + 1]) {

		if this.arr[startIndex * 2] > this.arr[startIndex * 2 + 1] {
			startIndex *= 2
		}else {
			startIndex = 2 * startIndex + 1
		}
	
	}
	fmt.Println("shit")
	return startIndex
}

func (this *MyHeap) PrintAll() {
	fmt.Println("haha")
	fmt.Println(this.arr[1:this.length+1])
}