package slice

import (
	"fmt"
	"math"
	"testing"

	"github.com/periaate/common"
)

var (
	i = timeI[int64]
	r = relTimeI[int64]
)

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
	{"2012", i(2012), true},
	{"2012/02", i(2012, 2), true},
	{"2012/02/13", i(2012, 2, 13), true},
	{"y2012", i(2012), true},
	{"y2012m2", i(2012, 2), true},
	{"y2012m2d13", i(2012, 2, 13), true},
	{"-y1", r(1), true},
	{"-m1", r(0, 1), true},
	{"-d7", r(0, 0, 7), true},
	{"-y1m3", r(1, 3), true},
	{"-y1m3d14", r(1, 3, 14), true},
}

var tos = []comb{
	{"", math.MaxInt64, true},
	{"2014", i(2014), true},
	{"2014/08", i(2014, 8), true},
	{"2014/08/22", i(2014, 8, 22), true},
	{"y2014", i(2014), true},
	{"y2014m8", i(2014, 8), true},
	{"y2014m8d22", i(2014, 8, 22), true},
	{"-y2", r(2), true},
	{"-2m", r(0, 2), true},
	{"-14d", r(0, 0, 14), true},
	{"-y1m3", r(1, 3), true},
	{"-y1m3d22", r(1, 3, 22), true},
}

func TestParseTimeSlice(t *testing.T) {
	combs := common.Combinations(froms, tos)
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
		{
			"t[-y4]",
			r(4),
			r(3), false},
		{
			"t[-m4]",
			r(0, 4),
			r(0, 3), false},
		{
			"t[-d4]",
			r(0, 0, 4),
			r(0, 0, 3), false},
		{
			"t[-y4m2]",
			r(4, 2),
			r(4, 1), false},
		{
			"t[-y4m2d4]",
			r(4, 2, 4),
			r(4, 2, 3), false},
	}
	for _, tt := range combs {
		r := fmt.Sprintf("t[%s:%s]", tt[0].inp, tt[1].inp)
		err := tt[0].inp == "" && tt[1].inp == ""
		cases = append(cases, timeExpects{r, tt[0].val, tt[1].val, err})
	}

	for _, tt := range cases {
		from, to, err := ParseTimeSlice(tt.pat)
		if tt.err {
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			return
		}
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if from != tt.from || to != tt.to {
			t.Errorf("not expected from %s, GOT:\n[%v, %v]\n[%v, %v]", tt.pat, from, to, tt.from, tt.to)
		}
	}
}
