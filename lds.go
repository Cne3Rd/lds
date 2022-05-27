package lds

import "fmt"

// lds is a simple package that allows
// you to create and work with list

type Node struct {
	Property interface{}
	NextNode *Node
}

type linkedList struct {
	headNode *Node
}

func NewList() *linkedList {
	list := &linkedList{}
	return list
}

// add node to the head of list
func (linkedList *linkedList) AddToHead(property interface{}) {
	node := Node{}
	node.Property = property

	if node.NextNode == nil {
		node.NextNode = linkedList.headNode
	}
	linkedList.headNode = &node
}

// itereate through a list
func (linkedList *linkedList) IterateList() {
	currentNode := linkedList.headNode
	for currentNode != nil {
		fmt.Println(currentNode.Property)
		currentNode = currentNode.NextNode
	}
}

// return the last node in the list
func (linkedList *linkedList) LastNode() *Node {
	currentNode := linkedList.headNode
	for currentNode != nil {
		if currentNode.NextNode == nil {
			lastNode := currentNode
			return lastNode
		}
		currentNode = currentNode.NextNode

	}
	return nil
}

// add node to the end of the list
func (linkedList *linkedList) AddToEnd(property interface{}) {
	node := Node{}
	node.Property = property

	if linkedList.headNode == nil {
		linkedList.headNode = &node
		return
	}

	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.NextNode = &node
	}

}

// return node having the value provided
// return first match node
func (linkedList *linkedList) NodeWithValue(property interface{}) *Node {
	node := linkedList.headNode
	for node != nil {
		if node.Property == property {
			return node
		}

		node = node.NextNode
	}

	return nil
}

// add node after a specific node
func (linkedList *linkedList) AddAfter(nodeProperty int, property interface{}) {
	nodewith := linkedList.NodeWithValue(nodeProperty)
	node := Node{}
	node.Property = property
	if nodewith != nil {
		node.NextNode = nodewith.NextNode
		nodewith.NextNode = &node
	}
}

// delete a specific node having the value provided
func (linkedList *linkedList) DeleteNode(property interface{}) {
	current := linkedList.headNode
	nodewith := linkedList.NodeWithValue(property)
	for current != nil {
		if linkedList.headNode.Property == property {
			linkedList.headNode = linkedList.headNode.NextNode
			return
		}
		if current.NextNode.Property == nodewith.Property {
			current.NextNode = nodewith.NextNode
			return

		}
		current = current.NextNode
	}

}
