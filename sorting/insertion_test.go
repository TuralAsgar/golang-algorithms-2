package sorting

import (
	"testing"
)

/* Commonly Used Slices */

var reverse = []int{20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1}
var sorted = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
var random = []int{14,3,8,11,6,13,19,15,5,20,9,18,12,4,16,1,17,10,7,2}

/* Tests Section */

func testInsertion(array []int, t * testing.T) {
    insertion_sort(array)
    for index, x := range array {
        if index > 0 && x < array[index -1] {
            t.FailNow()
        }
    }
}

func TestInsertionEmptyArray(t * testing.T) { testInsertion([]int{}, t) }
func TestInsertionArrayOfOne(t * testing.T) { testInsertion([]int{1}, t) }
func TestInsertionEmptyOfTwo(t * testing.T) { testInsertion([]int{2,1}, t) }
func TestInsertionReversedArray(t * testing.T) { testInsertion(reverse, t) }
func TestInsertionRandomArray(t * testing.T) { testInsertion(random, t) }
func TestInsertionSortedArray(t * testing.T) { testInsertion(sorted, t) }

/* Benchmark Section */

func benchmarkInsertion(array []int, b * testing.B) {
	for i := 0; i < b.N; i++ {
        insertion_sort(array)
    }
}

func BenchmarkReverseInsertion(b * testing.B) { benchmarkInsertion(reverse, b) }
func BenchmarkRandomInsertion(b * testing.B) { benchmarkInsertion(random, b) }
func BenchmarkSortedInsertion(b * testing.B) { benchmarkInsertion(sorted, b) }
