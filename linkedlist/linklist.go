package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
}

func main() {
	linkedlist := &LinkedList{}
	linkedlist.AddElement(5)
	linkedlist.AddElement(6)
	linkedlist.AddElement(7)
	linkedlist.PrintElement()
	linkedlist.SizeOfList()
	linkedlist.RemoveFirstElement()
	linkedlist.PrintElement()
	linkedlist.AddElement(9)
	linkedlist.AddElement(8)
	linkedlist.RemoveLastElement()
	linkedlist.PrintElement()

	linkedlist.SearchElement(9)
}

func (ll *LinkedList) AddElement(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (ll *LinkedList) PrintElement() {
	current := ll.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func (ll *LinkedList) SizeOfList() int {
	list := 0
	current := ll.head
	for current != nil {
		list++
		current = current.next
	}
	fmt.Println("Number of elements of linked list are: ", list)
	return list
}
func (ll *LinkedList) RemoveFirstElement() {
	if ll.head != nil {
		fmt.Println("First element removed is : ", ll.head.value)
		ll.head = ll.head.next
	}
}
func (ll *LinkedList) RemoveLastElement() {
	if ll.head == nil {
		fmt.Print("Removed last element :", ll.head.value)
		return
	}
	if ll.head.next == nil {
		fmt.Print("Removed last element :", ll.head.value)
		ll.head = nil
		return
	}
	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	fmt.Print("Removed last element :", ll.head.value)
	current.next = nil
}
func (ll *LinkedList) SearchElement(search int) bool {
	current := ll.head
	for current != nil {
		if current.value == search {
			fmt.Print("Search element found")
			return true
		}
		current = current.next
	}
	fmt.Print("Search element not found")
	return false
}
