package main

import (
	"fmt"
	"golang-datastruct/util/myarray"
)

func main() {
	// TestMyArray()

	// var v interface{}
	// v = 1
	// switch t := v.(type) {
	// case int:
	// 	fmt.Println("int")
	// case float64:
	// 	fmt.Println("float64")
	// //... etc
	// default:
	// 	_ = t
	// 	fmt.Println("unkown")
	// }
}

func TestMyArray() {
	arr := myarray.New()

	arr.data = make([]int, 4, 4)

	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}

	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst("a")
	fmt.Println(arr)

	arr.Remove(2)
	fmt.Println(arr)

	arr.RemoveElement(4)
	fmt.Println(arr)

	arr.RemoveFirst()
	fmt.Println(arr)

	for i := 0; i < 7; i++ {
		arr.RemoveFirst()
		fmt.Println(arr)
	}
}
