package wrs

import (
	"crypto/rand"
	"math/big"
	"sort"
)

func r(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}

// Choice represents a weighted value.
type Choice struct {
	W int         // weight
	V interface{} // value
}

// Choices is a slice of Choice elements.
type Choices []Choice

func (cs Choices) Len() int {
	return len(cs)
}

func (cs Choices) Less(i int, j int) bool {
	return cs[i].W < cs[j].W
}

func (cs Choices) Swap(i int, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// Choose returns a random element from cs.
//
// Total number of weights in cs should be >= 1.
func (cs Choices) Choose() interface{} {
	var totals []int
	runningTotal := 0
	for _, w := range cs {
		runningTotal += w.W
		totals = append(totals, runningTotal)
	}
	n := r(runningTotal) + 1
	i := sort.SearchInts(totals, n)
	return cs[i].V
}
