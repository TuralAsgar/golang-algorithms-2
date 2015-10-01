
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

var limit = 1001
var prng, csprng []int

func init() {
	prng = make([]int, limit)
	csprng = make([]int, limit)

	max := *big.NewInt(2)
	rand.Seed( time.Now().UTC().UnixNano() )
	for i:=0; i < limit; i++ {
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
	
	if ones / limit > 0 { fmt.Println(ones / limit); t.FailNow() }
	if ones / limit < 1 { fmt.Println(ones / limit); t.FailNow() }
}

func TestFrequencyPRNG(t * testing.T) { testFrequency(prng, t) }
func TestFrequencyCSPRNG(t * testing.T) { testFrequency(csprng, t) }
