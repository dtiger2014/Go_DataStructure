package algo_sort

import(
	"fmt"
	"math"
)

func MergeSort(arr []int, n int) {
	mergeSort(arr, 0, n-1)
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return;
	}

	mid := int(math.Ceil((l + r) / 2))
	fmt.Printf("merge Sort l=%v, mid=%v, r=%v\n", l, mid, r)

	
	// mergeSort(arr, l, mid)
	// mergeSort(arr, mid, r)
	// merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	// aux := make([]int, r-l+1)

	fmt.Printf("arr=%v, l=%v, mid=%v, r=%v\n", arr, l, mid, r)
	aux := arr[l:r+1]

	i := l // 0
	j := mid + 1 // 4
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-1]
			j++
		}
	}

} 