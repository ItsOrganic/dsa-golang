package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func main() {
	queue := &Queue{}
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(20)
	queue.PrintQueue()

	//Using dequeue
	fmt.Println("Eliminated element:", queue.Dequeue())
	queue.PrintQueue()

	//Checking if the queue is Empty
	fmt.Println("Queue is Empty:", queue.isEmpty())
	fmt.Println("first element of the queue:", queue.GetFirstElement())
	fmt.Println("Size of the queue:", queue.GetSize())

}

// Append the element at tbe last of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue(Remove from the first )
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]

	// Deleted items is stored in the item
	return item
}
func (q *Queue) PrintQueue() {
	fmt.Println(q.items)
}

func (q *Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return false
	}
	return true
}

func (q *Queue) GetFirstElement() int {
	if len(q.items) == 0 {
		return 0
	}
	return q.items[0]
}
func (q *Queue) GetSize() int {
	return len(q.items)
}
