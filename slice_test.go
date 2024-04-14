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

func sin() []int {
	return []int{0, 1, 2, 3, 4, 5}
}

var smallIn = []int{0, 1, 2, 3, 4, 5}
var shuffleAct = Shuffle[int]("#9")

func shuffledSin(in []int) []int {
	res, _ := shuffleAct([]int{}, in)
	return res
}

var expects = []expect{
	{arg: "[0]",
		in:  smallIn,
		out: []int{0}},
	{arg: "[3:]",
		in:  smallIn,
		out: []int{3, 4, 5}},
	{arg: "[:3]",
		in:  smallIn,
		out: []int{0, 1, 2}},
	{arg: "[3:5]",
		in:  smallIn,
		out: []int{3, 4}},
	{arg: "[2:-3]",
		in:  smallIn,
		out: []int{2}},
	{arg: "[-3:2]",
		in:  smallIn,
		out: []int{3, 4, 5, 0, 1}},
	{arg: "[=2]",
		in:  smallIn,
		out: []int{0, 1}},
	{arg: "[1=2]",
		in:  smallIn,
		out: []int{2, 3}},
	{arg: "[:2=2]",
		in:  smallIn,
		out: []int{0, 1, 2, 3}},
	{arg: "[//]",
		in:  smallIn,
		out: []int{5, 0, 1, 2, 3, 4}},
	{arg: "[\\]",
		in:  smallIn,
		out: []int{1, 2, 3, 4, 5, 0}},
	{arg: `[\\]`,
		in:  smallIn,
		out: []int{1, 2, 3, 4, 5, 0}},
	{arg: `[//2]`,
		in:  smallIn,
		out: []int{4, 5, 0, 1, 2, 3}},
	{arg: `[\\2]`,
		in:  smallIn,
		out: []int{2, 3, 4, 5, 0, 1}},
	{arg: `[//2.1:-2]`,
		in:  smallIn,
		out: []int{5, 0, 1}},
	{arg: `[1:4#9]`,
		in:  sin(),
		out: shuffledSin(sin()[1:4])},
	{arg: `[3:5-:]`,
		in:  smallIn,
		out: []int{4, 3}},
	{arg: `[2_-1]`,
		in:  smallIn,
		out: []int{0, 1, 3, 4, 5, 2}},
	{arg: `[-3_-1]`,
		in:  smallIn,
		out: []int{0, 1, 2, 4, 5, 3}},
	{arg: `[5_0]`,
		in:  smallIn,
		out: []int{5, 0, 1, 2, 3, 4}},
}

func TestSliceLang(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	// sar := []string{
	// 	"[1]",
	// 	"[3:5]",
	// 	"[-2:1]",
	// 	"[3=5]",
	// 	"[4=2//3]",
	// 	"[1:-1=2_5]",
	// 	"[1:-1=2_5 0]",
	// 	"[1:-1=2_5.1:2=5 0]",
	// }

	// for _, s := range sar {
	// 	a := Parse[string](s)
	// 	fmt.Println(a == nil)
	// }

	for _, e := range expects {
		t.Run(fmt.Sprint(e.arg), func(t *testing.T) {
			act := Parse[int](e.arg)
			out, err := act(e.in, e.in)
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
