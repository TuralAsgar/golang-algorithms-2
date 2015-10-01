package sorting

import (
     "floor"
)

func merge_sort(array []int) {
     if len(array) <= 1 { return }

     var left, right, middle
     middle = floor( len(array) / 2 )

     left = array[:middle]
     right = array[middle+1:]

     left = merge_sort(left)
     right = merge_sort(right)

     return merge(left, right)
}

func merge(left, right []int) {
     
}