package slice

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Clamp[T constraints.Ordered](val, lower, upper T) (res T) {
	if val >= upper {
		return upper
	}
	if val <= lower {
		return lower
	}
	return val
}

func SmartClamp[N constraints.Integer](a, l N) N { return ((a % l) + l) % l }

func ExtendedSlice[T any, N constraints.Integer](a, b, l N, inp []T) (res []T) {
	a = SmartClamp(a, l)
	b = SmartClamp(b, l)

	fmt.Println(a, b, l, len(inp), cap(inp))
	if a < b {
		return res[a:b]
	}
	return append(res[:a], res[:b]...)
}

func Slice[T any](pat string, inp []T) (res []T, err error) {
	switch {
	case len(pat) == 0:
		err = fmt.Errorf("input is empty")
		return
	case len(pat) > 3:
		if pat[0] == '[' && pat[len(pat)-1] == ']' {
			pat = pat[1 : len(pat)-1]
		}
	}
	var from, to int

	ind := strings.IndexRune(pat, ':')
	switch ind {
	case -1:
		from, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
		to = from + 1
	case 0:
		to, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
	case len(pat) - 1:
		to = len(inp)
		from, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
	default:
		from, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
		to, err = strconv.Atoi(pat)
		if err != nil {
			return
		}
	}

	return ExtendedSlice(from, to, len(inp), inp), nil
}

// func Sl[T any](pat string, input []T) (res []T, err error) {
// 	var from, to int

// 	if len(input) == 0 {
// 		err = fmt.Errorf("input is empty")
// 		slog.Error(err.Error())
// 		return
// 	}

// 	if len(pat) < 3 {
// 		err = fmt.Errorf("last argument is not long enough to be a slice")
// 		slog.Error(err.Error())
// 		return
// 	}

// 	L := len(pat) - 1
// 	if pat[0] != '[' || pat[L] != ']' {
// 		err = fmt.Errorf("last argument does not match slice pattern, does not start and end with brackets")
// 		slog.Error(err.Error())
// 		return
// 	}
// 	pat = pat[1:L]
// 	slog.Debug("slice pattern", "pattern", pat)

// 	for _, r := range pat {
// 		if !(r == '-' || r == '+' || r == ':' || r == pageToken || r >= '0' || r <= '9') {
// 			err = fmt.Errorf("slice pattern included non integer values")
// 			slog.Error(err.Error())
// 			return
// 		}
// 	}

// 	pageSize := 1

// 	if ind := strings.Index(pat, string(pageToken)); ind != -1 {
// 		vl := pat[ind+1:]
// 		pageSize, err = strconv.Atoi(vl)
// 		if err != nil {
// 			slog.Debug(err.Error())
// 			return
// 		}

// 		pat = pat[:ind]

// 		if ind == 0 {
// 			from = 0
// 			to = pageSize
// 			slog.Debug("page token only", "from", from, "to", to, "pagesize", pageSize, "input length", len(input))
// 			from = Clamp(from, 0, len(input))
// 			to = Clamp(to, 0, len(input))
// 			from = Clamp(from, 0, to)
// 			slog.Debug("clamped results", "from", from, "to", to, "pagesize", pageSize, "input length", len(input))
// 			return input[from:to], nil
// 		}
// 	}

// 	ind := strings.Index(pat, ":")

// 	if ind == -1 {
// 		slog.Debug("slice pattern is only one character long")
// 		if pat[0] == ':' {
// 			return input, nil
// 		}

// 		if pat[0] == '-' {
// 			slog.Debug("negative single index")
// 			from, err = parseMinus(pat, len(input), pageSize)
// 			if err != nil {
// 				slog.Debug(err.Error())
// 				return
// 			}
// 			to = (from + 1) * pageSize
// 		} else {
// 			from, err = strconv.Atoi(pat)
// 			if err != nil {
// 				slog.Debug(err.Error())
// 				return
// 			}

// 			to = from + 1
// 			to *= pageSize
// 			from *= pageSize
// 		}

// 		slog.Debug("slice results", "from", from, "to", to, "pagesize", pageSize, "input length", len(input))
// 		from = Clamp(from, 0, len(input))
// 		to = Clamp(to, 0, len(input))
// 		from = Clamp(from, 0, to)
// 		slog.Debug("clamped results", "from", from, "to", to, "pagesize", pageSize, "input length", len(input))
// 		return input[from:to], nil
// 	}

// 	ind = Clamp(ind, 0, len(pat))

// 	fromTo := []string{pat[:ind], pat[ind+1:]}

// 	if fromTo[f] == "" {
// 		from = 0
// 	} else {
// 		if fromTo[f][0] == '-' {
// 			from, err = parseMinus(fromTo[f], len(input), pageSize)
// 			if err != nil {
// 				slog.Error(err.Error())
// 				return
// 			}
// 		} else {
// 			from, err = strconv.Atoi(fromTo[f])
// 			if err != nil {
// 				slog.Error(err.Error())
// 				return
// 			}
// 			from *= pageSize
// 		}
// 	}

// 	if fromTo[t] == "" {
// 		to = len(input)
// 	} else {
// 		switch fromTo[t][0] {
// 		case '+':
// 			to, err = parsePlus(fromTo[t], from, pageSize)
// 			if err != nil {
// 				slog.Error(err.Error())
// 				return
// 			}
// 		case '-':
// 			to, err = parseMinus(fromTo[t], len(input), pageSize)
// 			if err != nil {
// 				slog.Error(err.Error())
// 				return
// 			}
// 		default:
// 			to, err = strconv.Atoi(fromTo[t])
// 			if err != nil {
// 				slog.Error(err.Error())
// 				return
// 			}
// 			to *= pageSize
// 		}
// 	}

// 	if from < 0 {
// 		from = len(input) - from
// 	}

// 	if to < 0 {
// 		to = len(input) - to
// 	}

// 	from = Clamp(from, 0, len(input))
// 	to = Clamp(to, 0, len(input))

// 	if from > to {
// 		res = input[from%len(input):]
// 		res = append(res, input[:to]...)
// 	} else {
// 		from = Clamp(from, 0, len(input))
// 		to = Clamp(to, 0, len(input))
// 		from = Clamp(from, 0, to)
// 		res = input[from:to]
// 	}

// 	return res, nil
// }

// func parseMinus(s string, l, size int) (int, error) {
// 	val, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return l + val*size, nil
// }

// func parsePlus(s string, l, size int) (int, error) {
// 	val, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return l + val*size, nil
// }
