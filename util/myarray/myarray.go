package myarray

const (
	// DefaultCap : 默认数组容量
	DefaultCap = 10
)

// ArrayData : 动态数组 struct
type ArrayData struct {
	size int
	data [DefaultCap]interface{}
}

func (arr *ArrayData) Init() {
	arr.size = 0
}
func (arr *ArrayData) GetCapacity() int {
	return len(arr.data)
}

func (arr *ArrayData) GetSize() int {
	return arr.size
}

func (arr *ArrayData) IsEmpty() bool {
	return arr.size == 0
}

func (arr *ArrayData) Add(index int, E interface{}) error {
	if index < 0 || index > arr.size {
		// return error.New("")
		return nil
	}

	// if size == len(arr.data){
	// 	arr.resize()
	// }
}

func (arr *ArrayData) resize(newCapacity int) {
	newArr := [newCapacity]interface{}
	for i, v := range arr.data {
		newArr[i] = v
	}
}