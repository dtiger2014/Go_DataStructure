package mystack

import (
	"fmt"
	"golang-datastruct/util/myarray"
)

const (
	// DefaultCap: 默认容量
	DefaultCap = 10
)

// ArrayStackData : array stack
type ArrayStackData struct {
	data *myarray.ArrayData
}

func New() *ArrayStackData {
	return &ArrayStackData{
		data: myarray.New(),
	}
}

func (stack *ArrayStackData) GetCapacity() int {
	return stack.data.GetCapacity()
}

func (stack *ArrayStackData) GetSize() int {
	return stack.data.GetSize()
}

func (stack *ArrayStackData) IsEmpty() bool {
	return stack.data.IsEmpty()
}

func (stack *ArrayStackData) Push(e interface{}) {
	stack.data.AddLast(e)
}

func (stack *ArrayStackData) Pop() (interface{}, error) {
	return stack.data.RemoveLast()
}

func (stack *ArrayStackData) Peek() (interface{}, error) {
	return stack.data.GetLast()
}

func (stack *ArrayStackData) String() string {
	str := fmt.Sprintf("Stack:\n")
	str += "["
	for i := 0; i < stack.data.GetSize(); i++ {
		v, _ := stack.data.Get(i)
		str += fmt.Sprintf("%v", v)
		if i != stack.data.GetSize()-1 {
			str += ", "
		}
	}
	str += "] top"

	return str
}
