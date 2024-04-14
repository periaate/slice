package slice

import (
	"testing"
)

type expect struct {
	args []string
	in   []int
	out  []int
}

var smallIn = []int{0, 1, 2, 3, 4, 5}

var expects = []expect{
	{args: []string{"[0]"},
		in:  smallIn,
		out: []int{0}},
	{args: []string{"[3:]"},
		in:  smallIn,
		out: []int{3, 4, 5}},
	{args: []string{"[:3]"},
		in:  smallIn,
		out: []int{0, 1, 2}},
	{args: []string{"[3:5]"},
		in:  smallIn,
		out: []int{3, 4}},
	{args: []string{"[2:-3]"},
		in:  smallIn,
		out: []int{2}},
	{args: []string{"[-3:2]"},
		in:  smallIn,
		out: []int{3, 4, 5, 0, 1}},
	{args: []string{"[=2]"},
		in:  smallIn,
		out: []int{0, 1}},
	{args: []string{"[1=2]"},
		in:  smallIn,
		out: []int{2, 3}},
	{args: []string{"[:2=2]"},
		in:  smallIn,
		out: []int{0, 1, 2, 3}},
	{args: []string{"[//]"},
		in:  smallIn,
		out: []int{5, 0, 1, 2, 3, 4}},
	{args: []string{"[\\]"},
		in:  smallIn,
		out: []int{1, 2, 3, 4, 5, 0}},
	{args: []string{`[\\]`},
		in:  smallIn,
		out: []int{1, 2, 3, 4, 5, 0}},
	{args: []string{`[//2]`},
		in:  smallIn,
		out: []int{4, 5, 0, 1, 2, 3}},
	{args: []string{`[\\2]`},
		in:  smallIn,
		out: []int{2, 3, 4, 5, 0, 1}},
	{args: []string{`[//2]`, `[1:3]`},
		in:  smallIn,
		out: []int{5, 0}},
	{args: []string{`[1:3//2]`},
		in:  smallIn,
		out: []int{0, 3, 4, 1, 2}},
}

func TestSlice(t *testing.T) {
	for r := range expects {
		_ = r
	}
}

/*
items inside `<>` are optional
items inside `[]` can be any size, also optional
`.` is a split operator

<->[0-9]<:><->[0-9]<=>[0-9]
{ // | \\ | \ | _ }[0-9]
*/
