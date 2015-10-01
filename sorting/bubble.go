package sorting

func bubble_sort(array []int) {
    swapped := true
    for swapped == true {
		swapped = false
		for index, x := range array {
			if index >= len(array) - 1 { break }
			if x > array[index + 1] {
				swap(array, index, index + 1)
				swapped = true
			} 
		}
	}
}
