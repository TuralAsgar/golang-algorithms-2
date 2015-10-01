package sorting

func selection_sort(array []int) {
     for index, _ := range array {
     	 min := index
     	 for i := index; i < len(array); i++ {
	     if array[min] > array[i] {
	     	min = i
	     }
	 }  
	 if min != index {
     	    swap(array, index, min)
     	 }
     }
}