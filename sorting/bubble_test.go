package sorting

import (
	"testing"
	"fmt"
)

/* Tests Section */

func testBubble(array []int, t * testing.T) {
    bubble_sort(array)
    for index, x := range array {
        if index > 0 && x < array[index -1] {
			fmt.Println("Out of order:", x, "<", array[index -1])
            t.FailNow()
        }
    }
}

func TestBubbleEmptyArray(t * testing.T) { testBubble([]int{}, t) }
func TestBubbleArrayOfOne(t * testing.T) { testBubble([]int{1}, t) }
func TestBubbleEmptyOfTwo(t * testing.T) { testBubble([]int{2,1}, t) }
func TestBubbleReversedArray(t * testing.T) { testBubble(reverse, t) }
func TestBubbleRandomArray(t * testing.T) { testBubble(random, t) }
func TestBubbleSortedArray(t * testing.T) { testBubble(sorted, t) }

/* Benchmark Section */

func benchmarkBubble(array []int, b * testing.B) {
	for i := 0; i < b.N; i++ {
        bubble_sort(array)
    }
}

func BenchmarkReverseBubble(b * testing.B) { benchmarkBubble(reverse, b) }
func BenchmarkRandomBubble(b * testing.B) { benchmarkBubble(random, b) }
func BenchmarkSortedBubble(b * testing.B) { benchmarkBubble(sorted, b) }
