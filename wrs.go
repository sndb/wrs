package wrs

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sort"
)

func randInt(max int) int {
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

// New returns a new Chooser or error when sum of weights is less than zero.
func New(cs ...Choice) (*Chooser, error) {
	c := new(Chooser)
	for _, cc := range cs {
		c.runningTotal += cc.W
		c.totals = append(c.totals, c.runningTotal)
	}
	if c.runningTotal < 1 {
		return nil, ErrSumOfWeights
	}
	c.Choices = cs
	return c, nil
}

// Pick returns a random element from Chooser.
func (c *Chooser) Pick() interface{} {
	x := randInt(c.runningTotal) + 1
	i := sort.SearchInts(c.totals, x)
	return c.Choices[i].V
}

func (c *Chooser) Len() int {
	return len(c.Choices)
}

func (c *Chooser) Less(i int, j int) bool {
	return c.Choices[i].W < c.Choices[j].W
}

func (c *Chooser) Swap(i int, j int) {
	c.Choices[i], c.Choices[j] = c.Choices[j], c.Choices[i]
}
