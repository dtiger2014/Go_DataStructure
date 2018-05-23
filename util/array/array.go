package util

import(
	"fmt"
)

type myArray struct {
	data []int
}

func MyArray(cap int) *myArray {
	return &myArray{
		data: make([]int, 0, cap),
	}
}

func (this *myArray) GetCapacity() int{
	return cap(this.data)
}

func (this *myArray) GetSize() int {
	return len(this.data)
}

func (this *myArray) IsEmpty() bool {
	return len(this.data) == 0
}

func (this *myArray) AddFirst(e int) {
	this.Add(0, e)
}

func (this *myArray) AddLast(e int) {
	this.Add(len(this.data), e)
}

func (this *myArray) Add(index int, e int) {
	if len(this.data) == cap(this.data) {
		panic("Add failed. Array is full")
	}
	
	if index < 0 || index > len(this.data) {
		panic("Add failed. Require index >= 0 and index <= size.")
	}
	
	// 保存 slice 中 index 后面的元素
	slice :=  append([]int{}, this.data[index:]...)
	this.data = append(append(this.data[:index], e), slice...)
}

func (this *myArray) Set(index int, e int) {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	this.data[index] = e
}

func (this *myArray) Get(index int) int {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

func (this *myArray) Contains(e int ) bool {
	for _, v := range this.data {
		if v == e {
			return true
		}
	}
	return false
}

func (this *myArray) Find(e int) int {
	for i, v := range this.data {
		if v == e {
			return i
		}
	}
	return -1
}

func (this *myArray) Remove(index int) int {
	if index < 0 || index >= len(this.data) {
		panic("Remove failed. Index is illegal")
	}

	ret := this.data[index]
	this.data = append(this.data[:index], this.data[index+1:]...)
	return ret
}

func (this *myArray) RemoveFirst() int {
	return this.Remove(0)
}

func (this *myArray) RemoveLast() int {
	return this.Remove(len(this.data) - 1)
}

func (this *myArray) RemoveElement(e int) {
	index := this.Find(e)
	if index != -1 {
		this.Remove(index)
	}
}

func (this *myArray) ToString() {
	fmt.Printf("Array: size = %v , capacity = %v\n", len(this.data), cap(this.data))
	fmt.Println(this.data)
}