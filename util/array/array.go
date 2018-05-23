package util

import(
	"fmt"
)

type MyArray struct {
	data [10]int
	size int
	cap int
}

func (arr *MyArray) GetCapacity() int{
	return arr.cap
}

func (arr *MyArray) GetSize() int {
	return arr.size
}

func (arr *MyArray) IsEmpty() bool {
	return arr.size == 0
}

func (arr *MyArray) AddFirst(e int) {
	arr.Add(0, e)
}

func (arr *MyArray) AddLast(e int) {
	arr.Add(arr.size, e)
}

func (arr *MyArray) Add(index int, e int) {
	if arr.size == len(arr.data) {
		panic("Add failed. Array is full")
	}
	
	if index < 0 || index > arr.size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i + 1] = arr.data[i]
	}

	arr.data[index] = e

	arr.size++
}

func (arr *MyArray) Get(index int) int {
	if index < 0 || index >= arr.size {
		panic("Get failed. Index is illegal.")
	}
	return arr.data[index]
}

func (arr *MyArray) Set(index int, e int) {
	if index < 0 || index >= arr.size {
		panic("Get failed. Index is illegal.")
	}
	arr.data[index] = e
}

func (arr *MyArray) ToString() {
	fmt.Println("Array: size = %d , capacity = %d\n", arr.size, len(arr.data))
	fmt.Println(arr.data)
}