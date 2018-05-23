package main

import(
	util "util/array"
	"strconv"
	"fmt"
	"time"
)

func main() {
	// myArrayInt()
	// myArrayString()
	// myArrayStudent()
	myArrayTest()

	/*
	result:
		RemoveLast() : O(1)
		RemoveFirst() : O(n)
		AddLast() : O(1)
		AddFirst() : O(n)

		Last 和 First 相差 700倍
	*/
}

func myArrayTest() {
	start := time.Now()
	// myarray ---
	array := util.MyArray(10)
	cnt := 10000
	for i := 0; i < cnt; i++ {
		array.AddLast(i)
		// array.AddFirst(i)
	}

	for i := 0; i < cnt; i++ {
		array.RemoveLast()
		// array.RemoveFirst()
	}
	array.ToString()
	// -----------
	useTime := time.Since(start)
	fmt.Println("Use time : ", useTime)
}

func myArrayInt() {
	array := util.MyArray(2)

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

func myArrayString() {
	array := util.MyArray(20)

	for i := 0; i < 10; i++ {
		array.AddLast(strconv.Itoa(i))
	}
	array.ToString()

	array.Add(1, "100")
	array.ToString()

	array.AddFirst("-1")
	array.ToString()

	array.Remove(2)
	array.ToString()

	array.RemoveElement("4")
	array.ToString()

	array.RemoveFirst()
	array.ToString()
}

type Student struct{
	name string
	score int
}
func myArrayStudent() {
	array := util.MyArray(20)
	array.AddLast(Student{name:"student1",score:100})
	array.AddLast(Student{name:"student2",score:80})
	array.AddLast(Student{name:"student3",score:70})
	array.ToString()
}