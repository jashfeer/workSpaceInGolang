package main

import "fmt"

type queue struct {
	items []int
}

func (q *queue) enQueue(n int) {
	q.items = append(q.items, n)
}

func (q *queue) deQueue() int {
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

func main() {
	myQueue := queue{}
	myQueue.enQueue(10)
	myQueue.enQueue(20)
	myQueue.enQueue(30)
	fmt.Println(myQueue)
	myQueue.deQueue()
	fmt.Println(myQueue)

}
