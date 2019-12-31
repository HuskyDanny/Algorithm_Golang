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
	this.percolateUp(val)

	return true
}

func (this *MyHeap) RemoveTop() int {
	
	topValue := this.arr[1]
	this.swap(1, this.length)
	this.length -= 1

	this.percolateDown(this.arr[1], 1)

	//this.swap(1, targetIndex)

	return topValue
}

func (this *MyHeap) percolateUp(val int) {

	startIndex := (this.length )

	for startIndex > 1  && this.arr[startIndex / 2] < val {
		parentIndex := startIndex
		startIndex = startIndex / 2

		this.swap(parentIndex, startIndex)
	}

	
	
}

func (this *MyHeap) swap(indexI int, indexJ int) {
	temp := this.arr[indexI]
	this.arr[indexI] = this.arr[indexJ]
	this.arr[indexJ] = temp
}

func (this *MyHeap) percolateDown(val int, startIndex int)  {
	

	for startIndex <= this.length && (startIndex * 2 <= this.length && val < this.arr[startIndex * 2] || (startIndex * 2 + 1 <= this.length) && val < this.arr[startIndex * 2 + 1]) {
		parentIndex := startIndex
		if (startIndex * 2 + 1 <= this.length) && this.arr[startIndex * 2] < this.arr[startIndex * 2 + 1] {
			startIndex = 2 * startIndex + 1
		}else {
			startIndex = 2 * startIndex
		}

		this.swap(parentIndex, startIndex)
	
	}
}

func (this *MyHeap) PrintAll() {
	fmt.Println(this.arr[1:this.length+1])
	
}

func (this *MyHeap) Heapify(arr []int) {
	for i := 1; i <= len(arr); i++ {
		this.arr[i] = arr[i-1]
		this.length++
	}

	for i := this.length / 2; i > 0; i-- {
		this.percolateDown(this.arr[i], i)
	}
}