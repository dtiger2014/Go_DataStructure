package main

import(
	sort "algo_sort"
	"fmt"
)

func main(){
	n := 20000;

	arr11 := sort.GenerateRandomArray(n, 0, n)
	arr12 := sort.CopyArray(arr11)
	arr13 := sort.CopyArray(arr11)
	arr14 := sort.CopyArray(arr11)
	arr15 := sort.CopyArray(arr11)
	arr16 := sort.CopyArray(arr11)
	arr17 := sort.CopyArray(arr11)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr11, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr12, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr13, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr14, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr15, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr16, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr17, n)
	
	fmt.Println()

	arr21 := sort.GenerateRandomArray(n, 0, 3)
	arr22 := sort.CopyArray(arr21)
	arr23 := sort.CopyArray(arr21)
	arr24 := sort.CopyArray(arr21)
	arr25 := sort.CopyArray(arr21)
	arr26 := sort.CopyArray(arr21)
	arr27 := sort.CopyArray(arr21)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr21, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr22, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr23, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr24, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr25, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr26, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr27, n)

	fmt.Println()

	arr31 := sort.GenerateNearlyOrderedArray(n, 100)
	arr32 := sort.CopyArray(arr31)
	arr33 := sort.CopyArray(arr31)
	arr34 := sort.CopyArray(arr31)
	arr35 := sort.CopyArray(arr31)
	arr36 := sort.CopyArray(arr31)
	arr37 := sort.CopyArray(arr31)

	sort.TestSort("SelectionSort", sort.SelectionSort, arr31, n)
	sort.TestSort("SelectionSort2", sort.SelectionSort2, arr32, n)
	sort.TestSort("InsertionSort", sort.InsertionSort, arr33, n)
	sort.TestSort("BubbleSort", sort.BubbleSort, arr34, n)
	sort.TestSort("BubbleSort2", sort.BubbleSort2, arr35, n)
	sort.TestSort("BubbleSort3", sort.BubbleSort3, arr36, n)
	sort.TestSort("ShellSort", sort.ShellSort, arr37, n)
}


