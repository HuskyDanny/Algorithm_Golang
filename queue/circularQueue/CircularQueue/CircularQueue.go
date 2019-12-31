package CircularQueue

import (
	"fmt"
)

type CircularQueue struct {
	buffer []interface{}
	head int 
	tail int
	capacity int
}

func NewCircularQueue(capacity int) *CircularQueue {
	if capacity == 0  {
		return nil
	}
	
	//initiate slice 
	return &CircularQueue{make([]interface{}, capacity), 0, 0, capacity}
}

func (this *CircularQueue) isEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

func (this *CircularQueue) isFull() bool {
	if (this.tail + 1) % this.capacity == this.head {
		return true
	}
	return false
}

func (this *CircularQueue) Enqueue(v interface{}) bool {
	if (this.isFull()) {
		return false
	}

	this.buffer[this.tail] = v
	this.tail = (this.tail + 1) % this.capacity

	return true
}

func (this *CircularQueue) Dequeue() interface{} {
	if (this.isEmpty()) {
		return nil
	}
	
	temp := this.buffer[this.head]
	this.head = (this.head + 1) % this.capacity
	return temp
}

func (this *CircularQueue) PrintAll() {

	for !this.isEmpty() {
		fmt.Println(this.Dequeue())
	}
}
