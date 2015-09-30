package sorting

func swap(array []int, i, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
