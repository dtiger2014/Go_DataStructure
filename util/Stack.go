package util

type stack struct {
	data *myArray
}

func Stack(cap int) *stack {
	return &stack{
		data: MyArray(cap),
	}
}

func (this *stack) GetSize() int {
	return this.data.GetSize()
}

func (this *stack) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *stack) GetCapacity() int {
	return this.data.GetCapacity()
}

func (this *stack) Push(e E) {
	this.data.AddLast(e)
}

func (this *stack) Pop() E {
	return this.data.RemoveLast()
}

func (this *stack) Peek() E {
	return this.data.GetLast()
}

func (this *stack) ToString() {
	this.data.ToString()
}