package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
	
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 60}
	node2 := &node{data: 70}
	node3 := &node{data: 80}
	node4 := &node{data: 90}
	node5 := &node{data: 50}
	node6 := &node{data: 40}
	node7 := &node{data: 960}
	node8 := &node{data: 590}
	node9 := &node{data: 790}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)
	mylist.prepend(node8)
	mylist.prepend(node9)
	mylist.printListData()
	mylist.deleteValue(90)
	mylist.printListData()
	mylist.deleteValue(790)
	mylist.printListData()
	mylist.deleteValue(90)
	mylist.printListData()

}
