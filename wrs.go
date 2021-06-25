package wrs

import (
	"crypto/rand"
	"errors"
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

// ErrSumOfWeights is returned by NewChooser when sum of weights is less than zero.
var ErrSumOfWeights = errors.New("sum of weights should be >= 1")

// Choice represents a weighted value.
type Choice struct {
	W int         // weight
	V interface{} // value
}

// Chooser represents a container of weighted choices to Pick from.
type Chooser struct {
	totals       []int
	runningTotal int
	Choices      []Choice
}

func (chr Chooser) Len() int {
	return len(chr.Choices)
}

func (chr Chooser) Less(i int, j int) bool {
	return chr.Choices[i].W < chr.Choices[j].W
}

func (chr Chooser) Swap(i int, j int) {
	chr.Choices[i], chr.Choices[j] = chr.Choices[j], chr.Choices[i]
}

// NewChooser returns a new Chooser or error when sum of weights is less than zero.
func NewChooser(cs ...Choice) (*Chooser, error) {
	chr := new(Chooser)
	for _, c := range cs {
		chr.runningTotal += c.W
		chr.totals = append(chr.totals, chr.runningTotal)
	}
	if chr.runningTotal < 1 {
		return nil, ErrSumOfWeights
	}
	chr.Choices = cs
	return chr, nil
}

// Pick returns a random element from Chooser.
func (chr *Chooser) Pick() interface{} {
	n := r(chr.runningTotal) + 1
	i := sort.SearchInts(chr.totals, n)
	return chr.Choices[i].V
}
