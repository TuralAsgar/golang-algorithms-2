package sorting

func insertion_sort(array []int) {
	for index, x := range array {
		for index > 0 && x < array[index - 1] {
			swap(array, index, index - 1)
			index--
		}
	}
}
