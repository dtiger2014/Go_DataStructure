package util

import(
	"fmt"
)

type node struct {
	e E
	next *node
}

func (this *node) toString() {
	fmt.Println(this.e)
}

type LinkedList struct {
	dummyHead *node
	size int
}

func (this *LinkedList) GetSize() int {
	return this.size
}

func (this *LinkedList) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedList) AddFirst(e E) {
	this.Add(0, e)
}

func (this *LinkedList) AddLast(e E) {
	this.Add(this.size, e)
}

func (this *LinkedList) Add(index int, e E) {
	if index < 0 || index > this.size {
		panic("Add failed. Illegal index.")
	}

	prev := this.dummyHead;
	for i := 0 ; i < index; i++ {
		prev = prev.next
	}

	// node := node{e, nil}
	// node.next = prev.next
	// prev.next = &node

	prev.next = &node{e, prev.next}   // Todo list
	this.size++
}

func (this *LinkedList) ToString() {
	fmt.Printf("LinkedList: size=%v [", this.size)

	cur := this.dummyHead
	for cur != nil {
		fmt.Printf("%v => ", cur.e)
		cur = cur.next
	}
	fmt.Println("]")
}
