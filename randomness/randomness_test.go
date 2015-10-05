
package randomness

import (
	"testing"
	"math/rand"
	"math/big"
	"fmt"
	"strconv"
	crand "crypto/rand"
	"time"
)

var limit = 1000
var prng, csprng []int

func init() {
	prng = make([]int, limit)
	csprng = make([]int, limit)

	max := *big.NewInt(2)
	rand.Seed( time.Now().UTC().UnixNano() )
	for i := 0; i < limit; i++ {
		prng[i] = rand.Intn(2)
		x, _ := crand.Int(crand.Reader, &max)
		csprng[i], _ = strconv.Atoi(x.String())
	}
}

func testFrequency(array []int, t * testing.T) {
	ones, zeros := 0, 0
	for _, x := range array {
		if x == 1 { ones++ } else { zeros++ }
	}

	avg := float32(ones / limit)
	if avg > float32(53 / 100) { fmt.Println("Avg:", avg); t.FailNow() }
	if avg < float32(47 / 100) { fmt.Println("Avg:", avg); t.FailNow() }
}

func testRuns(array []int, t * testing.T) {
	runSize := 25
	var runs = make([]int, runSize)
	var expected int
	var run = 1

	for i := 0; i < limit; i++ {
		if i == 0 {	expected = (array[i] + 1) % 2 } 

		for i < limit && expected == array[i] {
			run ++
			//fmt.Println(array[i])
			expected = (array[i] + 1) % 2
			i++
		}
		if run > 1 {
			runs[run] = runs[run] + 1
			run = 1
		} else { expected = (array[i] + 1) % 2 }
	}
	for i := 15; i < runSize; i++ {
		if runs[i] > 0 { fmt.Println(runs[i], "Run(s) with length", i); t.Fail()}
	}
}

func testSingleRuns(array []int, t * testing.T) {
	runSize := 35
	var runs = make([]int, runSize)
	var expected int
	var run = 0

	for i := 0; i < limit; i++ {
		if i == 0 { expected = array[i] }
		for i < limit && expected == array[i] {
			run ++
			i ++
		}
		if run > 1 {
			runs[run] = runs[run] + 1
			run = 0
		} else { expected = array[i] }
	}
    for i := 15; i < runSize; i++ {
        if runs[i] > 0 { fmt.Println(runs[i], "Single Run(s) with length", i); t.Fail() }
    }
}

func TestFrequencyPRNG(t * testing.T) { testFrequency(prng, t) }
func TestFrequencyCSPRNG(t * testing.T) { testFrequency(csprng, t) }
func TestRunsCSPRNG(t * testing.T) { testRuns(csprng, t) }
func TestRunsPRNG(t * testing.T) { testRuns(prng, t) }
func TestSingleRunsCSPRNG(t * testing.T) { testSingleRuns(csprng, t) }
func TestSingleRunsPRNG(t * testing.T) { testSingleRuns(prng, t) }
