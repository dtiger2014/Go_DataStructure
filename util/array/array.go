package util

import(
	"fmt"
)

type myArray struct {
	data []interface{}
}

func MyArray(cap int) *myArray {
	return &myArray{
		data: make([]interface{}, 0, cap),
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

func (this *myArray) AddFirst(e interface{}) {
	this.Add(0, e)
}

func (this *myArray) AddLast(e interface{}) {
	this.Add(len(this.data), e)
}

func (this *myArray) Add(index int, e interface{}) {
	if index < 0 || index > len(this.data) {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if len(this.data) == cap(this.data) {
		// panic("Add failed. Array is full")
		this.resize(2 * cap(this.data))
	}
	
	// 保存 slice 中 index 后面的元素
	slice := append([]interface{}{}, this.data[index:]...)
	this.data = append(append(this.data[:index], e), slice...)
}

func (this *myArray) Set(index int, e interface{}) {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	this.data[index] = e
}

func (this *myArray) Get(index int) interface{} {
	if index < 0 || index >= len(this.data) {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

func (this *myArray) Contains(e interface{} ) bool {
	for _, v := range this.data {
		if v == e {
			return true
		}
	}
	return false
}

func (this *myArray) Find(e interface{}) int {
	for i, v := range this.data {
		if v == e {
			return i
		}
	}
	return -1
}

func (this *myArray) Remove(index int) interface{} {
	if index < 0 || index >= len(this.data) {
		panic("Remove failed. Index is illegal")
	}

	ret := this.data[index]
	this.data = append(this.data[:index], this.data[index+1:]...)

	if (len(this.data) == cap(this.data) / 4 && len(this.data) / 2 != 0) {
		this.resize(cap(this.data) / 2)
	}

	return ret
}

func (this *myArray) RemoveFirst() interface{} {
	return this.Remove(0)
}

func (this *myArray) RemoveLast() interface{} {
	return this.Remove(len(this.data) - 1)
}

func (this *myArray) RemoveElement(e interface{}) {
	index := this.Find(e)
	if index != -1 {
		this.Remove(index)
	}
}

func (this *myArray) resize(newCapacity int) {
	slice := this.data
	this.data = make([]interface{}, 0, newCapacity)
	this.data = append(this.data, slice...)
}

func (this *myArray) ToString() {
	fmt.Printf("Array: size = %v , capacity = %v\n", len(this.data), cap(this.data))
	fmt.Println(this.data)
}