package sorting

import (
     "math"
)

func merge_sort(array []int) []int {
    if len(array) <= 1 { return array }

    var middle int
	var left, right []int

    middle = int(math.Floor( float64(len(array) / 2) ))

    left = array[:middle]
    right = array[middle+1:]

    left = merge_sort(left)
    right = merge_sort(right)

    return merge(left, right)
}

func merge(left, right []int) []int {
	sizeL := len(left)
	sizeR := len(right)
	if sizeL >= 1 && sizeR == 0 { return left }
	if sizeL == 0 && sizeR >= 1 { return right }

	i, j := 0, 0
	var result []int
	for i < sizeL && j < sizeR {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < sizeL {
		result = append(result, left[i])
		i++
	}
	for j < sizeR {
		result = append(result, right[j])
		j++
	}
	return result
}
