package slice

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

/*

items inside `<>` are optional
items inside `[]` can be any size, also optional
`.` is a split operator
int = [0-9]
val = <->int
sel = <val><:><val><=>int
shf = <sel>{ // | \\ | \ | _ }int

[0]
[3:]
[:3]
[3:5]
[2:-3]
[-3:2]
[=2]
[1=2]
[:2=2]



[0-9]		= values
`-` 		= `len`  `relative` `negative`	= `[-1]` `[:-10]`
`=`			= `value` size operator			= `[3=5]` `[-2=10]`


`//` = shift right	[0, 1, 2] [//]	-> [2, 0, 1]
`\\` = shift left	[0, 1, 2] [\\]	-> [1, 2, 0]
	[1, 2, 3, 4, 5] [//2]			-> [4, 5, 1, 2, 3]
	[1, 2, 3, 4, 5] [\\2]			-> [3, 4, 5, 1, 2]
	[1, 2, 3, 4, 5] [//2] [1:-2]	-> [5, 1]
	[1, 2, 3, 4, 5] [:2//2]			-> [3, 4, 1, 2, 5]



`_`	= shift into	<src>_<dst> if either undefined, implicit 0
	[1, 2, 3, 4, 5] [-1_2]	-> [1, 2, 5, 3, 4]
-3_5


// `.`		= combine operator (deduplicates) on-hold
`?!` 	= shuffle operator (randomizes)
` `		= then operator

:
=
//
\\
_
.
?!
*/

/*
<operation>
<function>
<combine>
<...>
*/

/*
map args -> matches
reduce arr -> match(curr) -> curr

split . -> res
map res -> act
reduce act arr -> arr
*/

type Act[T any] func(arr []T, sub []T) (res []T, err error)

func slice[T any](arg string) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		return Slice("["+arg+"]", sub)
	}
}

func shiftRight[T any](arg string) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		return ShiftRight[T](arg, arr, sub)
	}
}

func shiftLeft[T any](arg string) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		return ShiftLeft[T](arg, arr, sub)
	}
}

func shiftInto[T any](arg string) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		return ShiftInto[T](arg, arr, sub)
	}
}

/*
operations
	slices

functions
	shifts
*/

func Parse[T any](str string) (act Act[T]) {
	l := len(str) - 1
	if str[0] != '[' && str[l] != ']' {
		return
	}
	str = str[1:l]

	shuffle := strings.Contains(str, "?!")
	if shuffle {
		return
	}

	partsU := strings.FieldsFunc(str, func(r rune) bool { return r == ' ' })
	parts := make([][]string, 0, len(partsU))
	for _, part := range partsU {
		parts = append(parts, SplitWithAll(part,
			`//`,
			`\`,
			`\\`,
			`_`,
		))
	}

	acts := make([]Act[T], 0)

	for _, part := range parts {
		for _, s := range part {
			switch {
			case strings.Contains(s, `//`):
				acts = append(acts, shiftRight[T](s))
			case strings.Contains(s, `\`):
				acts = append(acts, shiftLeft[T](s))
			case strings.Contains(s, `_`):
				acts = append(acts, shiftInto[T](s))
			default:
				acts = append(acts, slice[T](s))
			}
		}
	}

	return ToAct(acts)
}

func ToAct[T any](fns []Act[T]) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		res = sub[:]
		for _, fn := range fns {
			res, err = fn(sub, res)
		}
		return
	}
}

func ShiftRight[T any](arg string, arr []T, sub []T) (res []T, err error) {
	n := shiftParse(arg)

	return shiftDirection(arr, sub, n, true)
}
func ShiftLeft[T any](arg string, arr []T, sub []T) (res []T, err error) {
	n := shiftParse(arg)

	return shiftDirection(arr, sub, n, false)
}
func ShiftInto[T any](arg string, arr []T, sub []T) (res []T, err error) {
	dst := shiftParse(arg)

	from := len(arr) - cap(sub)
	to := from + len(sub)
	res = arr

	if from == to {
		return
	}
	if dst >= to && dst < from {
		err = fmt.Errorf("destination is within the source")
		return
	}
	if dst > to {
		dst -= len(sub)
	}

	dst = Clamp(dst, 0, len(sub))

	res = append(res[:dst], append(sub, res[dst:]...)...)

	return
}

func shiftParse(arg string) int {
	arg = strings.TrimLeft(arg, `/`)
	arg = strings.TrimLeft(arg, `\`)
	if len(arg) == 0 {
		return 1
	}
	n, err := strconv.Atoi(arg)
	if err != nil {
		return 0
	}
	return n
}

func shiftDirection[T any](arr []T, sub []T, by int, R bool) (res []T, err error) {
	res = sub
	slog.Debug("moving all")
	if R {
		nr := make([]T, len(sub))
		for i := 0; i < len(sub); i++ {
			j := (i + by) % len(sub)
			slog.Debug("moving", "src", i, "dst", j)
			nr[j] = sub[i]
		}
		res = nr
	} else {
		nr := make([]T, len(sub))
		for i := 0; i < len(sub); i++ {
			src := i
			dst := (len(sub) - by + i) % len(sub)
			slog.Debug("moving", "src", src, "dst", dst)
			nr[dst] = sub[src]
		}
		res = nr
	}

	return
}
