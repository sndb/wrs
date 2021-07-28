package wrs

import (
	"fmt"
	"testing"
)

func TestWRS(t *testing.T) {
	cases := []struct {
		in   []Choice
		want error
	}{
		{
			[]Choice{{5, "a"}, {3, "b"}, {2, "c"}}, nil,
		},
		{
			[]Choice{{1, "x"}, {0, "y"}, {1, "z"}}, nil,
		},
		{
			[]Choice{{0, "i"}, {1, "j"}}, nil,
		},
		{
			[]Choice{{0, "z"}, {0, "x"}}, ErrSumOfWeights,
		},
	}
	for i, cs := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c, err := New(cs.in...)
			if err != cs.want {
				t.Fatal(err)
			}
			if err != nil {
				return
			}
			m := make(map[string]int)
			for i := 0; i < 10000; i++ {
				m[fmt.Sprintf(c.Pick().(string))]++
			}
			for k, v := range m {
				fmt.Println(k, v)
			}
		})
	}
}
