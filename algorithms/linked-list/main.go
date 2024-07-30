package LinkedList

import (
	"errors"
	"fmt"
)

type Node struct {
	Data     string
	NextNode *Node
}

type LinkedList struct {
	First *Node
	Last  *Node
}

var (
	ErrNodeNotFound  = errors.New("++ Node not found")
	ErrEmptyList     = errors.New("++ Error: Impossible to add a node after another node in an empty list")
	ErrNodeNotInList = errors.New("++ Error: Node not found in the list")
)

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

func CreateNode(data string) *Node {
	return &Node{Data: data, NextNode: nil}
}

func (list *LinkedList) AddNode(n *Node) {
	if list.First == nil {
		// List is empty, this is the first and last AddNode
		list.First = n
		list.Last = n
	} else {
		// The last node should point to the new node
		list.Last.NextNode = n

		// while the new node is now the last
		list.Last = n
	}
}

func (list *LinkedList) FindNodeByData(data string) (*Node, error) {
	var current = list.First

	for current != nil {
		if current.Data == data {
			return current, nil
		}

		current = current.NextNode
	}

	return nil, ErrNodeNotFound
}

func (list *LinkedList) FindNodeByAddress(n *Node) (*Node, error) {
	var current = list.First

	for current != nil {
		if current == n {
			return current, nil
		}
		current = current.NextNode
	}

	return nil, ErrNodeNotFound
}

func (list *LinkedList) AddNodeAfter(newNode *Node, after *Node) error {

	if list.First == nil {
		return ErrEmptyList
	}

	_, err := list.FindNodeByData(after.Data)

	if err != nil {
		return err
	}

	// Insert newNode after the found node
	newNode.NextNode = after.NextNode
	after.NextNode = newNode

	// Update the Last pointer if we inserted at the end
	if after.NextNode == list.Last {
		list.Last = newNode
	}

	// Added successfully
	return nil
}

func (list *LinkedList) PrintList() {
	var current = list.First

	fmt.Println("++ Current state of the list: ")
	for current != nil {
		fmt.Printf("%s\n", current.Data)
		current = current.NextNode
	}
	fmt.Println("")
}

func (list *LinkedList) ClearList() {
	list.First = nil
	list.Last = nil
}

func (list *LinkedList) DeleteNode(n *Node) error {
	if list.First == nil {
		return ErrEmptyList
	}

	if n == list.First {
		list.First = n.NextNode

		if n == list.Last {
			list.Last = nil
		}

		n.NextNode = nil
		return nil
	}

	// n is in the middle of the linked list
	// Find the node that -> to (n)
	previous := list.First
	for previous.NextNode != nil && previous.NextNode != n {
		previous = previous.NextNode
	}

	if previous.NextNode == nil {
		return ErrNodeNotInList
	}

	previous.NextNode = n.NextNode

	if n == list.Last {
		list.Last = previous
	}

	n.NextNode = nil

	return nil
}

func ExampleOne() {
	var linkedList = NewLinkedList()

	var firstNode = CreateNode("Jonathan")
	var secondNode = CreateNode("Luíza")
	var thirdNode = CreateNode("Julinha")

	linkedList.AddNode(firstNode)
	linkedList.AddNode(secondNode)
	linkedList.AddNode(thirdNode)

	linkedList.PrintList()
}

func FindExample() {
	var linkedList = NewLinkedList()

	var firstNode = CreateNode("Jonathan")
	var secondNode = CreateNode("Luíza")

	linkedList.AddNode(firstNode)
	linkedList.AddNode(secondNode)

	var n, error = linkedList.FindNodeByData("Luíza")

	if error == nil {
		fmt.Printf("++ Found node: %s\n\n", n.Data)
	}
}

func InsertAfterExample() {
	var linkedList = NewLinkedList()

	var firstNode = CreateNode("Jonathan")
	var secondNode = CreateNode("Luíza")
	var thirdNode = CreateNode("Julinha")
	var lastNode = CreateNode("Tia Marcela")

	linkedList.AddNode(firstNode)
	linkedList.AddNode(secondNode)
	linkedList.AddNode(thirdNode)

	// Add the last node after the second
	errorToInsertAfter := linkedList.AddNodeAfter(lastNode, secondNode)

	if errorToInsertAfter != nil {
		fmt.Printf("++ Error while trying to insert after: %e", errorToInsertAfter)
	}

	linkedList.PrintList()
}

func InsertEmptyExample() {
	// Try to insert after with an empty list:
	newLinkedList := NewLinkedList()

	var firstNode = CreateNode("Jonathan")
	var secondNode = CreateNode("Luíza")

	emptyInsertAfterError := newLinkedList.AddNodeAfter(firstNode, secondNode)

	fmt.Println(emptyInsertAfterError)
}

func DeleteExample() {
	var linkedList = NewLinkedList()

	var firstNode = CreateNode("Jonathan")
	var secondNode = CreateNode("Luíza")
	var thirdNode = CreateNode("Julinha")

	linkedList.AddNode(firstNode)
	linkedList.AddNode(secondNode)
	linkedList.AddNode(thirdNode)

	linkedList.DeleteNode(firstNode)
	linkedList.PrintList()
}
