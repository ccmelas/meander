package meander_test

import (
	"testing"

	"github.com/ccmelas/meander"
	"github.com/cheekybits/is"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)
	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}

func TestCostString(t *testing.T) {
	is := is.New(t)
	is.Equal(meander.Cost1.String(), "$")
	is.Equal(meander.Cost2.String(), "$$")
	is.Equal(meander.Cost3.String(), "$$$")
	is.Equal(meander.Cost4.String(), "$$$$")
	is.Equal(meander.Cost5.String(), "$$$$$")
}

func TestParseCost(t *testing.T) {
	is := is.New(t)
	is.Equal(meander.Cost1, meander.ParseCost("$"))
	is.Equal(meander.Cost2, meander.ParseCost("$$"))
	is.Equal(meander.Cost3, meander.ParseCost("$$$"))
	is.Equal(meander.Cost4, meander.ParseCost("$$$$"))
	is.Equal(meander.Cost5, meander.ParseCost("$$$$$"))
}
