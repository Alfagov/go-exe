package linkedlist

import "errors"

type Elm struct {
	data int
	next *Elm
}

type List struct {
	head *Elm
	size int
}

func New(inputNum []int) *List {
	output := &List{}

	for _, number := range inputNum {
		output.Push(number)
	}

	return output
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {

	l.head = &Elm{element, l.head}
	l.size++

}

func (l *List) Pop() (int, error) {

	// Check if list is empty
	if l.size < 1 {
		return 0, errors.New("no elements")
	}

	rmHead := l.head
	l.head = rmHead.next
	rmHead.next = nil
	l.size--
	return rmHead.data, nil
}

func (l *List) Array() []int {
	// Create an array of size l.size
	output := make([]int, l.size)

	// Iterate through the list and add the data to the array
	for i, head := l.size-1, l.head; i > -1; i, head = i-1, head.next {
		output[i] = head.data
	}

	return output
}

func (l *List) Reverse() *List {

	output := &List{}

	for head := l.head; head != nil; head = head.next {
		output.Push(head.data)
	}

	return output
}
