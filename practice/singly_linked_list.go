package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data     int32
	nextNode *Node
}

type List struct {
	head *Node
}

func (self *List) addNewNode(newNode *Node) error {
	if newNode == nil {
		return errors.New("node is empty...")
	}

	tempNextNode := self.head
	for tempNextNode.nextNode != nil {
		tempNextNode = tempNextNode.nextNode
	}

	tempNextNode.nextNode = newNode

	return nil
}

func (self *List) printList() error {
	if self.head == nil {
		return errors.New("head is empty...")
	}

	tempNextNode := self.head
	fmt.Printf("|> [head] ")
	for tempNextNode != nil {
		fmt.Printf("-> [data: %d]", tempNextNode.data)
		tempNextNode = tempNextNode.nextNode
	}

	fmt.Printf("\n")

	return nil
}
