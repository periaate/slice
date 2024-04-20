package slice

import (
	"fmt"
	"log/slog"
	"testing"
)

type expect struct {
	arg string
	in  []int
	out []int
}

func getArr() []int {
	return []int{0, 1, 2, 3, 4, 5}
}

const seed = "#9"

func getShuffledArr(in []int) []int {
	res, _ := shuffle(seed, in)
	return res
}

var expects = []expect{
	{arg: "[//]",
		in:  getArr(),
		out: []int{5, 0, 1, 2, 3, 4}},
	{arg: "[//.0]",
		in:  getArr(),
		out: []int{5}},
	{arg: "[//.:-1]",
		in:  getArr(),
		out: []int{5, 0, 1, 2, 3}},
	{arg: `[\\]`,
		in:  getArr(),
		out: []int{1, 2, 3, 4, 5, 0}},
	{arg: `[//2]`,
		in:  getArr(),
		out: []int{4, 5, 0, 1, 2, 3}},
	{arg: `[\\2]`,
		in:  getArr(),
		out: []int{2, 3, 4, 5, 0, 1}},
	{arg: `[//2.1:-2]`,
		in:  getArr(),
		out: []int{5, 0, 1}},
	{arg: `[1:4#9]`,
		in:  getArr(),
		out: getShuffledArr(getArr()[1:4])},
	// {arg: `[1:4][0]`,
	// 	in:  getArr(),
	// 	out: []int{1}},
	// {arg: `[:5][0_2_4]`,
	// 	in:  getArr(),
	// 	out: []int{0, 2, 4}},
}

func TestSliceLang(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	for _, e := range expects {
		t.Run(fmt.Sprint(e.arg), func(t *testing.T) {
			expr := NewExpression[int]()
			expr.Parse(e.arg)
			out, err := expr.Eval(e.in)
			if err != nil {
				t.Error(err)
			}
			if len(out) != len(e.out) {
				t.Errorf("expected %v, got %v", e.out, out)
			}
			for i := range out {
				if out[i] != e.out[i] {
					t.Errorf("expected %v, got %v", e.out, out)
				}
			}
		})
	}
}
