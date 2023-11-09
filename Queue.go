package main

type Queue struct {
	head   *queueNode
	tail   *queueNode
	length int
}

type queueNode struct {
	value int
	next  *queueNode
}

func createQueueNode(item int) *queueNode {
	return &queueNode{item, nil}
}

func (q *Queue) Enqueue(item int) {
	newNode := createQueueNode(item)
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}
