package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

func createNode(val int) *Node {
	return &Node{val, nil, nil}
}

func (ll *LinkedList) print() {
	curr := ll.head
	for {
		if curr == nil {
			fmt.Println("")
			return
		} else {
			fmt.Printf("%d, ", curr.val)
			curr = curr.next
		}
	}
}

func (ll *LinkedList) append(item int) {
	newNode := createNode(item)
	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		newNode.prev = ll.tail
		ll.tail = newNode
	}
	ll.length++
}

func (ll *LinkedList) prepend(item int) {
	newNode := &Node{item, nil, nil}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
		newNode.next.prev = newNode
	}
	ll.length++
}

func (ll *LinkedList) get(index int) (int, error) {
	currNode := ll.head
	if index > ll.length || index < 0 {
		return -1, errors.New("Index out of Bounds")
	}
	for x := 0; x <= index; x++ {
		currNode = currNode.next
	}
	return currNode.val, nil
}
func (ll *LinkedList) insertAt(item int, index int) {
	if index > ll.length || index < 0 {
		return
	}
	if index == 0 {
		ll.prepend(item)
		return
	}
	if index == ll.length {
		ll.append(item)
		return
	}
	newNode := &Node{item, nil, nil}
	currNode := ll.head
	for x := 0; x < index; x++ {
		currNode = currNode.next
	}
	newNode.next = currNode.next
	newNode.prev = currNode
	newNode.next.prev = newNode
	currNode.next = newNode
}
func (ll *LinkedList) remove(item int) (int, error) {
	currNode := ll.head
	for {
		if currNode == nil {
			break
		} else if currNode.val == item {
			currNode.prev.next = currNode.next
			currNode.next.prev = currNode.prev
			return currNode.val, nil
		} else {
			currNode = currNode.next
		}
	}
	return -1, errors.New("Not found")
}
func (ll *LinkedList) removeAt(index int) (int, error) {
	currNode := ll.head
	if index > ll.length || index < 0 {
		return -1, errors.New("Index out of Bounds")
	}
	for x := 0; x < index; x++ {
		currNode = currNode.next
	}
	currNode.prev.next = currNode.next
	currNode.next.prev = currNode.prev
	return currNode.val, nil
}
