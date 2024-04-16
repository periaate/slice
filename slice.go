package slice

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	split     = `.`
	shiftR    = `//`
	shiftL    = `\\`
	shiftLAlt = `\`
	shiftTo   = `_`
	shuffle   = `#`
	reverse   = `-:`
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

	partsU := strings.Split(str, split)
	parts := make([][]string, 0, len(partsU))
	for _, part := range partsU {
		parts = append(parts, SplitWithAll(part,
			shiftR,
			shiftL,
			shiftLAlt,
			shiftTo,
			shuffle,
			reverse,
		))
	}

	acts := make([]Act[T], 0)

	for _, part := range parts {
		for _, s := range part {
			switch {
			case strings.Contains(s, shiftR):
				acts = append(acts, shiftRight[T](s))
			case strings.Contains(s, shiftLAlt):
				acts = append(acts, shiftLeft[T](s))
			case strings.Contains(s, shiftTo):
				acts = append(acts, shiftInto[T](s))
			case strings.Contains(s, shuffle):
				acts = append(acts, Shuffle[T](s))
			case strings.Contains(s, reverse):
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

func randSeed(seed int64) *rand.Rand {
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	source := rand.NewSource(seed)
	return rand.New(source)
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
	src := randSeed(seed)
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
	var dst int
	if len(arg[1:]) != 0 {
		dst = shiftParse(arg[1:])
	}

	dst = (len(arr) + dst) % len(arr)
	start := len(arr) - cap(sub)
	end := start + len(sub)

	if dst > start && dst < end {
		err = fmt.Errorf("destination is within itself")
		slog.Error("error in shift into", "error", err)
		return
	}

	switch {
	case dst == len(arr)-1:
		res = append(arr, sub...)
	case dst == 0:
		res = append(sub, arr...)
		start += len(sub)
		end += len(sub)
	case dst > end:
		res = append(arr[:dst], append(sub, arr[dst:]...)...)
	case dst < start:
		res = append(arr[:dst], sub...)
		res = append(res, arr[dst:]...)

		start += len(sub)
		end += len(sub)
	}

	res = append(res[:start], res[end:]...)

	return
}

// func ShiftInto[T any](arg string, arr []T, sub []T) (res []T, err error) {
// 	dst := shiftParse(arg[1:])
// 	if dst < 0 {
// 		dst = len(arr) + dst
// 	}
// 	fmt.Println(dst, len(arr), len(sub))

// 	from := len(arr) - (cap(arr) - cap(sub))
// 	to := from + len(sub)
// 	res = arr

// 	if from == to {
// 		return
// 	}
// 	if dst >= to && dst < from {
// 		err = fmt.Errorf("destination is within the source")
// 		return
// 	}

// 	dst = dst % len(arr)

// 	res = append(res[:to], append(sub, res[dst:]...)...)

// 	return
// }

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
