package wrs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPick(t *testing.T) {
	cases := []struct {
		in               []Choice
		wantErr          error
		wantDistribution map[string]int
	}{
		{
			[]Choice{{5, "a"}, {3, "b"}, {2, "c"}},
			nil,
			map[string]int{"a": 5, "b": 3, "c": 2},
		},
		{
			[]Choice{{1, "x"}, {0, "y"}, {1, "z"}},
			nil,
			map[string]int{"x": 5, "z": 5},
		},
		{
			[]Choice{{0, "i"}, {1, "j"}},
			nil,
			map[string]int{"j": 10},
		},
		{
			[]Choice{{0, "z"}, {0, "x"}},
			ErrSumOfWeights,
			nil,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.in), func(t *testing.T) {
			chooser, err := New(c.in...)
			if err != c.wantErr {
				t.Errorf("got %v, want %v", err, c.wantErr)
			}
			if err != nil {
				return
			}

			m := make(map[string]int)
			for i := 0; i < 10000; i++ {
				m[chooser.Pick().(string)]++
			}
			for k, v := range m {
				m[k] = (v + 100) / 1000
			}

			if !reflect.DeepEqual(c.wantDistribution, m) {
				t.Errorf("want %v to be deeply equal to %v",
					m, c.wantDistribution)
			}
		})
	}
}
