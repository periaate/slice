package slice

import (
	"fmt"
	"testing"
)

// type expect struct {
// 	args []string
// 	in   []int
// 	out  []int
// }

// var smallIn = []int{0, 1, 2, 3, 4, 5}

// var expects = []expect{
// 	{args: []string{"[0]"},
// 		in:  smallIn,
// 		out: []int{0}},
// 	{args: []string{"[3:]"},
// 		in:  smallIn,
// 		out: []int{3, 4, 5}},
// 	{args: []string{"[:3]"},
// 		in:  smallIn,
// 		out: []int{0, 1, 2}},
// 	{args: []string{"[3:5]"},
// 		in:  smallIn,
// 		out: []int{3, 4}},
// 	{args: []string{"[2:-3]"},
// 		in:  smallIn,
// 		out: []int{2}},
// 	{args: []string{"[-3:2]"},
// 		in:  smallIn,
// 		out: []int{3, 4, 5, 0, 1}},
// 	{args: []string{"[=2]"},
// 		in:  smallIn,
// 		out: []int{0, 1}},
// 	{args: []string{"[1=2]"},
// 		in:  smallIn,
// 		out: []int{2, 3}},
// 	{args: []string{"[:2=2]"},
// 		in:  smallIn,
// 		out: []int{0, 1, 2, 3}},
// 	{args: []string{"[//]"},
// 		in:  smallIn,
// 		out: []int{5, 0, 1, 2, 3, 4}},
// 	{args: []string{"[\\]"},
// 		in:  smallIn,
// 		out: []int{1, 2, 3, 4, 5, 0}},
// 	{args: []string{`[\\]`},
// 		in:  smallIn,
// 		out: []int{1, 2, 3, 4, 5, 0}},
// 	{args: []string{`[//2]`},
// 		in:  smallIn,
// 		out: []int{4, 5, 0, 1, 2, 3}},
// 	{args: []string{`[\\2]`},
// 		in:  smallIn,
// 		out: []int{2, 3, 4, 5, 0, 1}},
// 	{args: []string{`[//2]`, `[1:3]`},
// 		in:  smallIn,
// 		out: []int{5, 0}},
// 	{args: []string{`[1:3//2]`},
// 		in:  smallIn,
// 		out: []int{0, 3, 4, 1, 2}},
// }

func TestSlice(t *testing.T) {
	sar := []string{
		"[1]",
		"[3:5]",
		"[-2:1]",
		"[3=5]",
		"[4=2//3]",
		"[1:-1=2_5]",
		"[1:-1=2_5 0]",
		"[1:-1=2_5.1:2=5 0]",
	}

	for _, s := range sar {
		a := Parse[string](s)
		fmt.Println(a == nil)
	}
}
