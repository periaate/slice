package slice

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	split     = `.`
	shiftR    = `//`
	shiftL    = `\\`
	randomize = `#`
)

type Act[T any] func(pat string, arr []T) (res []T, err error)
type actArgs[T any] struct {
	Arg string
	Fn  Act[T]
}

type Expression[T any] struct {
	Syntax map[string]Act[T]
	acts   []actArgs[T]
}

func (e *Expression[T]) Eval(arr []T) (res []T, err error) {
	for _, fn := range e.acts {
		arr, err = fn.Fn(fn.Arg, arr)
	}
	return arr, err
}

func NewExpression[T any]() *Expression[T] {
	syn := map[string]Act[T]{
		shiftR:    ShiftRight[T],
		shiftL:    ShiftLeft[T],
		randomize: shuffle[T],
	}
	return &Expression[T]{
		Syntax: syn,
		acts:   make([]actArgs[T], 0),
	}
}

func (e *Expression[T]) Parse(pat string) {
	l := len(pat) - 1
	if pat[0] == '[' && pat[l] == ']' {
		pat = pat[1:l]
	}

	partsU := strings.Split(pat, split)
	parts := make([][]string, 0, len(partsU))

	keys := make([]string, 0, len(e.Syntax))
	for k := range e.Syntax {
		keys = append(keys, k)
	}

	for _, part := range partsU {
		parts = append(parts, SplitWithAll(part, keys...))
	}

	for _, part := range parts {
		for _, pat := range part {
			fn, ok := e.Find(pat)
			switch {
			case ok:
				e.acts = append(e.acts, actArgs[T]{pat, fn})
			default:
				e.acts = append(e.acts, actArgs[T]{pat,
					func(pat string, arr []T) (res []T, err error) {
						return FromPattern(pat, arr)
					}})
			}
		}
	}
}

func (e *Expression[T]) Find(pat string) (Act[T], bool) {
	for k, v := range e.Syntax {
		if strings.Contains(pat, k) {
			return v, true
		}
	}
	return nil, false
}

func shuffle[T any](pat string, arr []T) (res []T, err error) {
	var seed int64
	if len(pat) > 0 {
		pat = pat[1:]
		seed, err = strconv.ParseInt(pat, 10, 64)
		if err != nil {
			seed = 0
		}
	}
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	src := rand.NewSource(seed)

	res = append(res, arr...)
	for i := range res {
		j := src.Int63() % int64(len(res))
		res[i], res[j] = res[j], res[i]
	}
	return res, nil
}

func ShiftRight[T any](arg string, arr []T) (res []T, err error) {
	var n int
	arg = strings.ReplaceAll(arg, `/`, "")

	if len(arg) > 0 {
		n, err = strconv.Atoi(arg)
		if err != nil {
			return arr, err
		}
	}
	return shiftDirection(arr, n, true)
}
func ShiftLeft[T any](arg string, arr []T) (res []T, err error) {
	var n int
	arg = strings.ReplaceAll(arg, `\`, "")

	if len(arg) > 0 {
		n, err = strconv.Atoi(arg)
		if err != nil {
			return arr, err
		}
	}
	return shiftDirection(arr, n, false)
}

func shiftDirection[T any](arr []T, by int, R bool) (res []T, err error) {
	if by < 1 {
		by = 1
	}
	res = arr
	if R {
		nr := make([]T, len(arr))
		for i := 0; i < len(arr); i++ {
			j := (i + by) % len(arr)
			nr[j] = arr[i]
		}
		res = nr
	} else {
		nr := make([]T, len(arr))
		for i := 0; i < len(arr); i++ {
			src := i
			dst := (len(arr) - by + i) % len(arr)
			if dst < 0 {
				dst += len(arr)
			}
			nr[dst] = arr[src]
		}
		res = nr
	}

	return
}
