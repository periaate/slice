package slice

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
)

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

func Parse[T any](str string) (act Act[T]) {
	l := len(str) - 1
	if str[0] != '[' && str[l] != ']' {
		return
	}
	str = str[1:l]

	partsU := strings.FieldsFunc(str, func(r rune) bool { return r == '?' })
	parts := make([][]string, 0, len(partsU))
	for _, part := range partsU {
		parts = append(parts, SplitWithAll(part,
			`//`,
			`\`,
			`\\`,
			`_`,
			`#`,
			`-:`,
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
			case strings.Contains(s, `#`):
				acts = append(acts, Shuffle[T](s))
			case strings.Contains(s, `-:`):
				acts = append(acts, Reverse[T])
			default:
				acts = append(acts, slice[T](s))
			}
		}
	}

	return ToAct(acts)
}

func Reverse[T any](_ []T, sub []T) (res []T, err error) {
	res = make([]T, 0, len(sub))
	for i := range sub {
		res = append(res, sub[len(sub)-1-i])
	}
	return
}

func Shuffle[T any](arg string) Act[T] {
	var seed int64
	var err error
	if len(arg) > 0 {
		arg = arg[1:]
		seed, err = strconv.ParseInt(arg, 10, 64)
		if err != nil {
			seed = 0
		}
	}
	src := rand.NewSource(seed)
	return func(arr []T, sub []T) (res []T, err error) {
		res = append(res, sub...)
		for i := range res {
			j := src.Int63() % int64(len(res))
			res[i], res[j] = res[j], res[i]
		}
		return res, nil
	}
}

func ToAct[T any](fns []Act[T]) Act[T] {
	return func(arr []T, sub []T) (res []T, err error) {
		res = make([]T, len(sub))
		copy(res, sub)
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

func shiftDirection[T any](_ []T, sub []T, by int, R bool) (res []T, err error) {
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
			if dst < 0 {
				dst += len(sub)
			}
			nr[dst] = sub[src]
		}
		res = nr
	}

	return
}
