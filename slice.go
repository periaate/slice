package slice

import (
	"fmt"
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


`.`		= combine operator (deduplicates)
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

type Act[T any] func(arr []T) (res []T, ok bool)

func Slice[T any](arr []T) (res []T, ok bool) {
	// [-2:1] negative direction (loop through start to end)

	return
}

func ShiftRight[T any](arr []T, from, to int) []T { return arr }
func ShiftLeft[T any](arr []T, from, to int) []T  { return arr }
func ShiftInto[T any](arr []T, from, to int) []T  { return arr }

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

	sar := SplitWithAll(str,
		`//`,
		`\`,
		`\\`,
		`_`,
		` `,
	)

	for _, s := range sar {
		isSlice := strings.ContainsAny(s, sliceIdents)
		isShift := strings.ContainsAny(s, shiftIdents)
		switch {
		case isShift:
			fmt.Println(s)
		case isSlice:
			fallthrough
		default:
			fmt.Println(s)
		}
	}

	return
}

const (
	shiftIdents = `/\_`
	sliceIdents = `:=`
)

func ToAct[T any](fns []Act[T]) Act[T] {
	return func(arr []T) (res []T, ok bool) {
		res = arr[:]
		for _, fn := range fns {
			res, ok = fn(res)
		}
		return
	}
}
