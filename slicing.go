package slice

import (
	"fmt"
	"strconv"
	"strings"

	c "github.com/periaate/common"
)

func Slice[T any](from, to, l int, inp []T) (res []T) {
	if from == to {
		return
	}
	// fromC := from
	// toC := to
	from = c.SmartClamp(from, l)
	to = c.SmartClamp(to, l)

	if from > to {
		return inp[to:from]
	}

	// fmt.Printf("FROM:	%v -> %v\n", fromC, from)
	// fmt.Printf("TO:	%v -> %v\n", toC, to)
	return inp[from:to]
}

func FromPattern[T any](pat string, inp []T) (res []T, err error) {
	if len(pat) == 0 {
		err = fmt.Errorf("input is empty")
		return inp, err
	}
	if len(pat) > 2 {
		if pat[0] == '[' && pat[len(pat)-1] == ']' {
			pat = pat[1 : len(pat)-1]
		}
	}
	if strings.Count(pat, ":") > 1 {
		err = fmt.Errorf("pattern contains too many ':'")
		return
	}
	if pat[0] == ':' && len(pat) == 1 {
		return inp, nil
	}

	var from, to int

	ind := strings.IndexRune(pat, ':')
	switch ind {
	case -1:
		from, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
		return append(res, inp[c.SmartClamp(from, len(inp))]), nil
	case 0:
		to, err = strconv.Atoi(pat[1:])
		if err != nil {
			return
		}
	case len(pat) - 1:
		to = len(inp)
		from, err = strconv.Atoi(pat[:len(pat)-1])
		if err != nil {
			return
		}
	default:
		pats := strings.Split(pat, ":")
		from, err = strconv.Atoi(pats[0])
		if err != nil {
			return
		}
		to, err = strconv.Atoi(pats[1])
		if err != nil {
			return
		}
	}
	// fmt.Printf("PAT:	%s\n", pat)
	return Slice(from, to, len(inp), inp), nil
}
