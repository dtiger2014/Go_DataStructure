package myarray

import (
	"errors"
	"fmt"
)

const (
	// DefaultCap : 默认数组容量
	DefaultCap = 10
)

// ArrayData : 动态数组 struct
type ArrayData struct {
	data []interface{}
	size int
}

func New() *ArrayData {
	return &ArrayData{
		data: make([]interface{}, DefaultCap, DefaultCap),
		size: 0,
	}
}

func (arr *ArrayData) GetCapacity() int {
	return cap(arr.data)
}

func (arr *ArrayData) GetSize() int {
	return arr.size
}

func (arr *ArrayData) IsEmpty() bool {
	return arr.size == 0
}

func (arr *ArrayData) Add(index int, e interface{}) error {
	if index < 0 || index > arr.size {
		return errors.New("Add failed. Require index >= 0 and index <= size.")
	}

	if arr.size == cap(arr.data) {
		arr.resize(2 * cap(arr.data))
	}

	for i := arr.size - 1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}

	arr.data[index] = e
	arr.size++

	return nil
}

func (arr *ArrayData) AddLast(e interface{}) {
	arr.Add(arr.size, e)
}

func (arr *ArrayData) AddFirst(e interface{}) {
	arr.Add(0, e)
}

func (arr *ArrayData) Get(index int) (interface{}, error) {
	if index < 0 || index > arr.size {
		return nil, errors.New("Get failed. Index is illegal.")
	}

	return arr.data[index], nil
}

func (arr *ArrayData) Set(index int, e interface{}) error {
	if index < 0 || index > arr.size {
		return errors.New("Set failed. Index is illegal.")
	}

	arr.data[index] = e

	return nil
}

func (arr *ArrayData) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

func (arr *ArrayData) Find(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

func (arr *ArrayData) Remove(index int) (interface{}, error) {
	if index < 0 || index >= arr.size {
		return nil, errors.New("Remove failed. Index is illegal.")
	}

	ret := arr.data[index]

	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	arr.data[arr.size] = nil

	if arr.size == cap(arr.data)/4 && cap(arr.data)/2 != 0 {
		arr.resize(cap(arr.data) / 2)
	}

	return ret, nil
}

func (arr *ArrayData) RemoveFirst() (interface{}, error) {
	return arr.Remove(0)
}

func (arr *ArrayData) RemoveLast() (interface{}, error) {
	return arr.Remove(len(arr.data))
}

func (arr *ArrayData) RemoveElement(e interface{}) {
	index := arr.Find(e)
	if index != 1 {
		_, _ = arr.Remove(index)
	}
}

func (arr *ArrayData) String() string {
	str := fmt.Sprintf("Array: size = %d , capacity = %d\n", arr.size, cap(arr.data))
	str += "["
	for i := 0; i < arr.size; i++ {
		str += fmt.Sprintf("%v", arr.data[i])
		if i != arr.size-1 {
			str += ", "
		}
	}
	str += "]"

	return str
}

func (arr *ArrayData) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity, newCapacity)
	for i := 0; i < arr.size; i++ {
		newArr[i] = arr.data[i]
	}

	arr.data = newArr
}
