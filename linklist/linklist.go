package linklist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type SinglyLinkList[T comparable] struct {
	head *Node[T]
}

func CreateSinglyLinkList[T comparable]() *SinglyLinkList[T] {
	return &SinglyLinkList[T]{head: nil}
}

func (s *SinglyLinkList[T]) Insert(data T) {

	node := &Node[T]{data: data, next: nil}

	if s.head == nil {
		s.head = node
		return
	}

	visitNode := s.head
	for visitNode.next != nil {
		visitNode = visitNode.next
	}
	visitNode.next = node
}

func (s *SinglyLinkList[T]) DisplayLinkList() {

	if s.head == nil {
		fmt.Println("Empty link list")
		return
	}

	for visitNode := s.head; visitNode != nil; visitNode = visitNode.next {
		fmt.Printf("%v-->", visitNode.data)
	}

	fmt.Println("nil")
}

func (s *SinglyLinkList[T]) Delete(data T) error {

	if s.head == nil {
		fmt.Println("Empty link list")
		return errors.New("Empty Link list")
	}

	current, prev := s.head, s.head

	for ; current != nil; current = current.next {

		if current.data == data {
			if prev == current { //this holds good only for one node or first node to be deleted
				s.head = current.next
				return nil
			}
			prev.next = current.next
			current.next = nil
			return nil
		}
		prev = current
	}

	return errors.New("Element not found")
}
