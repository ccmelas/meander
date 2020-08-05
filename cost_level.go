package meander

// Cost represents users budget
type Cost int8

// golint-disable-next-line
const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (c Cost) String() string {
	for k, v := range costStrings {
		if c == v {
			return k
		}
	}
	return "invalid"
}

//ParseCost returns the Cost value of a string
func ParseCost(s string) Cost {
	return costStrings[s]
}
