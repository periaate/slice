package slice

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/periaate/common"
	"golang.org/x/exp/constraints"
)

const (
	hour = time.Hour / 1000 / 1000 / 1000
	yDur = hour * 24 * 365
	y    = int64(yDur)
	mDur = hour * 24 * 365 / 12
	m    = int64(mDur)
	dDur = hour * 24 * 365
	d    = int64(dDur)
)

// can be provided {from date}:{to date}
// also accepts y{value} m{value} d{value} to specify just
// year, month, day, or some combination of them.
// if a [y|m|d] expression has "-" precedengin it, it is instead
// deducted from the current time.
// t[-y1:] would mean "in the last year until now"
// t[-2m:-1m] would mean "during the time between 60 and 30 days ago"
func ParseTimeSlice(pat string) (from, to int64, err error) {
	if len(pat) > 3 {
		if pat[:2] == "t[" && pat[len(pat)-1] == ']' {
			pat = pat[2 : len(pat)-1]
		}
	}
	to = math.MaxInt64

	if len(pat) < 2 {
		err = fmt.Errorf("pattern is too short to be a time slice")
		return
	}

	ind := strings.Index(pat, ":")
	spl := strings.Split(pat, ":")

	switch ind {
	case -1:
		from, to = Parse(pat)
	case len(pat) - 1: // to any
		from, _ = Parse(spl[0])
	case 0: // from any
		to, _ = Parse(pat)
	default:
		from, _ = Parse(spl[0])
		to, _ = Parse(spl[1])
	}
	return
}

func Parse(pat string, base ...int64) (unixT, off int64) {
	ng := pat[0] == '-'
	if ng {
		pat = pat[1:]
	}
	spl := common.SplitWithAll(pat, false, "/", "y", "m", "d")
	cnt := strings.Count(pat, "/")

	switch {
	case cnt == 2 || strings.Contains(pat, "d"):
		off = d
	case cnt == 1 || strings.Contains(pat, "m"):
		off = m
	case cnt == 0 || strings.Contains(pat, "y"):
		off = y
	}

	unixT = TimeI(Map(Int[int64], spl)...)

	if ng {
		if len(base) == 0 {
			base = append(base, time.Now().Unix())
		}
		unixT = base[0] - unixT
	}
	return unixT, unixT + off
}

func Map[A any, B any](f func(A) B, arr []A) []B {
	res := make([]B, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}

func Int[I constraints.Integer](pat string) I {
	i, err := strconv.ParseInt(pat, 10, 64)
	if err != nil {
		return I(0)
	}
	return I(i)
}

type Timeable interface {
	~int | ~int32 | ~int64 | ~uint32 | ~uint64
}

func TimeI[N Timeable](vals ...N) N {
	var unixT N
	for i, v := range vals {
		switch i {
		case 0:
			unixT += (v - 1970) * N(y)
		case 1:
			unixT += v * N(m)
		case 2:
			unixT += v * N(d)
		default:
			return unixT
		}
	}
	return unixT
}
