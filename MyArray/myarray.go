package main

import(
	util "util/array"
)

func main() {
	array := util.MyArray(20)

	for i := 0; i < 10; i++ {
		array.AddLast(i)
	}
	array.ToString()

	array.Add(1, 100)
	array.ToString()

	array.AddFirst(-1)
	array.ToString()

	array.Remove(2)
	array.ToString()

	array.RemoveElement(4)
	array.ToString()

	array.RemoveFirst()
	array.ToString()
}
