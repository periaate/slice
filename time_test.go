package slice

import (
	"fmt"
	"math"
	"testing"
	"time"
)

const fromDef int64 = 0

var now = time.Now()

func Y(t time.Time) int64 { return int64(t.Year()) }
func M(t time.Time) int64 { return int64(t.Month()) }
func D(t time.Time) int64 { return int64(t.Day()) }

type timeExpects struct {
	pat      string
	from, to int64
	err      bool
}

type comb struct {
	inp string
	val int64
	set bool
}

var froms = []comb{
	{"", 0, true},
	{"2012", TimeI[int64](2012), true},
	{"2012/02", TimeI[int64](2012, 2), true},
	{"2012/02/13", TimeI[int64](2012, 2, 13), true},
	{"y2012", TimeI[int64](2012), true},
	{"y2012m2", TimeI[int64](2012, 2), true},
	{"y2012m2d13", TimeI[int64](2012, 2, 13), true},
	{"-y1", TimeI(Y(now) - 1), true},
	{"-m1", TimeI(M(now) - 1), true},
	{"-d7", TimeI(D(now) - 7), true},
	{"-y1m3", TimeI(Y(now)-1, M(now)-3), true},
	{"-y1m3d14", TimeI(Y(now)-1, M(now)-3, D(now)-14), true},
}

var tos = []comb{
	{"", math.MaxInt64, true},
	{"2014", TimeI[int64](2014), true},
	{"2014/08", TimeI[int64](2014, 8), true},
	{"2014/08/22", TimeI[int64](2014, 8, 22), true},
	{"y2014", TimeI[int64](2014), true},
	{"y2014m8", TimeI[int64](2014, 8), true},
	{"y2014m8d22", TimeI[int64](2014, 8, 22), true},
	{"-y2", TimeI(Y(now) - 2), true},
	{"-2m", TimeI(M(now) - 2), true},
	{"-14d", TimeI(D(now) - 14), true},
	{"-y1m3", TimeI(Y(now)-1, M(now)-3), true},
	{"-y1m3d22", TimeI(Y(now)-1, M(now)-3, D(now)-22), true},
}

func TestParseTimeSlice(t *testing.T) {
	i := TimeI[int64]
	// combs := common.Combinations(froms, tos)
	cases := []timeExpects{
		{
			"t[2012]",
			i(2012),
			i(2013), false},
		{
			"t[2012/02]",
			i(2012, 2),
			i(2012, 3), false},
		{
			"t[2012/02/13]",
			i(2012, 2, 13),
			i(2012, 2, 14), false},
		{
			"t[y2012]",
			i(2012),
			i(2013), false},
		{
			"t[y2012m2]",
			i(2012, 2),
			i(2012, 3), false},
		{
			"t[y2012m2d13]",
			i(2012, 2, 13),
			i(2012, 2, 14), false},
		// {"t[-y4]", now.Unix() -
		// i(Y(now)-4), now.Unix() -
		// i(Y(now)-3), false},
		// 		// {"t[-m4]", now.Unix() -
		// i(M(now)-4), now.Unix() - i(M(now)-3), false},
		// 		// {"t[-d4]", now.Unix() -
		// i(D(now)-4), now.Unix() -
		// i(D(now)-3), false},
		// 		// {"t[-y4m2]", now.Unix() -
		// i(Y(now)-4, M(now)-2), now.Unix() -
		// i(Y(now)-4, M(now)-1), false},
		// 		// {"t[-y4m2d4]", now.Unix() -
		// i(Y(now)-4, M(now)-2, D(now)-4), now.Unix() -
		// i(Y(now)-4, M(now)-2, D(now)-3), false},
	}
	// fmt.Println(len(combs))
	// for _, tt := range combs {
	// 	r := fmt.Sprintf("t[%s:%s]", tt[0].inp, tt[1].inp)
	// 	cases = append(cases, timeExpects{r, tt[0].val, tt[1].val, false})
	// }
	fmt.Println(len(cases))

	for _, tt := range cases {
		from, to, err := ParseTimeSlice(tt.pat)
		if tt.err {
			if err == nil {
				t.Errorf("expected error, got nil")
			}
		}
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if from != tt.from || to != tt.to {
			t.Errorf("not expected from %s, GOT:\n[%v, %v]\n[%v, %v]", tt.pat, from, to, tt.from, tt.to)
		}
	}
}
