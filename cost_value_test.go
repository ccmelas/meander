package meander_test

import (
	"testing"

	"github.com/ccmelas/meander"
	"github.com/cheekybits/is"
)

func TestParseCostRange(t *testing.T) {
	is := is.New(t)
	var l meander.CostRange
	var err error
	l, err = meander.ParseCostRange("$$...$$$")
	is.NoErr(err)
	is.Equal(l.From, meander.Cost2)
	is.Equal(l.To, meander.Cost3)
	l, err = meander.ParseCostRange("$...$$$$$")
	is.NoErr(err)
	is.Equal(l.From, meander.Cost1)
	is.Equal(l.To, meander.Cost5)
}

func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	r := meander.CostRange{
		From: meander.Cost2,
		To:   meander.Cost4,
	}
	is.Equal("$$...$$$$", r.String())
}
