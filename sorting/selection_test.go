package sorting

import (
	"testing"
)

/* Tests Section */

func testSelection(array []int, t * testing.T) {
    selection_sort(array)
    for index, x := range array {
        if index > 0 && x < array[index -1] {
            t.FailNow()
        }
    }
}

func TestSelectionEmptyArray(t * testing.T) { testSelection([]int{}, t) }
func TestSelectionArrayOfOne(t * testing.T) { testSelection([]int{1}, t) }
func TestSelectionEmptyOfTwo(t * testing.T) { testSelection([]int{2,1}, t) }
func TestSelectionReversedArray(t * testing.T) { testSelection(reverse, t) }
func TestSelectionRandomArray(t * testing.T) { testSelection(random, t) }
func TestSelectionSortedArray(t * testing.T) { testSelection(sorted, t) }

/* Benchmark Section */

func benchmarkSelection(array []int, b * testing.B) {
	for i := 0; i < b.N; i++ {
        selection_sort(array)
    }
}

func BenchmarkReverseSelection(b * testing.B) { benchmarkSelection(reverse, b) }
func BenchmarkRandomSelection(b * testing.B) { benchmarkSelection(random, b) }
func BenchmarkSortedSelection(b * testing.B) { benchmarkSelection(sorted, b) }
