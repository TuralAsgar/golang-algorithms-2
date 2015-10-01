package sorting

import (
	"testing"
	"fmt"
)

/* Tests Section */

func testMerge(array []int, t * testing.T) {
    array = merge_sort(array)
    for index, x := range array {
        if index > 0 && x < array[index -1] {
			fmt.Println("Out of order:", x, "<", array[index -1])
            t.FailNow()
        }
    }
}

func TestMergeEmptyArray(t * testing.T) { testMerge([]int{}, t) }
func TestMergeArrayOfOne(t * testing.T) { testMerge([]int{1}, t) }
func TestMergeEmptyOfTwo(t * testing.T) { testMerge([]int{2,1}, t) }
func TestMergeReversedArray(t * testing.T) { testMerge(reverse, t) }
func TestMergeRandomArray(t * testing.T) { testMerge(random, t) }
func TestMergeSortedArray(t * testing.T) { testMerge(sorted, t) }

/* Benchmark Section */

func benchmarkMerge(array []int, b * testing.B) {
	for i := 0; i < b.N; i++ {
        merge_sort(array)
    }
}

func BenchmarkReverseMerge(b * testing.B) { benchmarkMerge(reverse, b) }
func BenchmarkRandomMerge(b * testing.B) { benchmarkMerge(random, b) }
func BenchmarkSortedMerge(b * testing.B) { benchmarkMerge(sorted, b) }
