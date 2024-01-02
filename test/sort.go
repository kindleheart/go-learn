package test_demo

func Sort(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	QuickSort(newArr, 0, len(newArr)-1)
	return newArr
}

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := arr[start+(end-start)>>1]
	l, r := start, end
	for l <= r {
		for l <= r && arr[l] < pivot {
			l++
		}
		for l <= r && arr[r] > pivot {
			r--
		}
		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	QuickSort(arr, start, r)
	QuickSort(arr, l, end)
}
