package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type queue struct {
	front *Node
	back  *Node
	size  int
}

func main() {
	myQueue := &queue{nil, nil, 0}
	fmt.Println("Created queue")
	fmt.Println()
	fmt.Print("Enqueueing items to the queue...\n\n")
	myQueue.enqueue("Mary")
	myQueue.enqueue("Jane")
	myQueue.enqueue("Xander")

	myQueue.printAllNodes()
	fmt.Println("---line break---")
	fmt.Println("Dequeue happening")
	fmt.Println("---line break---")
	myQueue.dequeue()
	myQueue.printAllNodes()

}

func (p *queue) dequeue() (string, error) {
	var item string

	if p.front == nil {
		return "", errors.New("Empty Queue!")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil

}

func (p *queue) enqueue(name string) error {
	newNode := &Node{name, nil}

	if p.front == nil { //if the queue is empty, add to front
		p.front = newNode
	} else { //after the first addition
		p.back.next = newNode //make current back struct point to new node
	}

	//*this is the tricky part
	//when p.front is nil
	//p.back also becomes the node1
	//when p.front is not nil
	//new node becomes back
	//p.back.next - node1.next becomes nwenode.(refer to line 62)
	p.back = newNode
	p.size++
	return nil
}

func (p *queue) printAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}

	fmt.Printf("%+v\n", currentNode.item)

	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}
