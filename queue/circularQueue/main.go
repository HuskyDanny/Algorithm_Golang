package main

import (
	cq "./CircularQueue"
)

func main() {
	q := cq.NewCircularQueue(5)

	q.Enqueue(2)
	q.Enqueue(5)
	q.Enqueue(1.212)
	q.Enqueue(5)
	q.Dequeue()
	q.Dequeue()
	q.PrintAll()
}