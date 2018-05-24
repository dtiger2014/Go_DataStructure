package functions

import(
	util "util"
	"fmt"
	// "time"
)

func TestLinkedList() {
	linkedList := util.LinkedList{}

	linkedList.AddFirst(1)
	// for i := 0; i < 5; i++ {
	// 	linkedList.AddLast(i)
	// 	// linkedList.ToString()
	// }
	fmt.Println(linkedList.GetSize())
}