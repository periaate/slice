package slice

import "strings"

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

type act struct {
}

func Slice[T any](r act, arr []T) (res []T, ok bool) {
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

func Parse(str string) {
	l := len(str) - 1
	if str[0] != '[' && str[l] != ']' {
		return
	}
	str = str[1:l]

	shuffle := strings.Contains(str, "?!")
	if shuffle {
		return
	}

	sar := strings.FieldsFunc(str, func(r rune) bool {
		return r == '/' || r == '\\' || r == '_' || r == '.'
	})

}
