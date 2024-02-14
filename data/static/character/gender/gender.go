package gender

//go:generate stringer -type=Gender
type Gender int

func (g Gender) Count() int {
	return 6
}

const (
	Undefined Gender = iota
	Intersex
	Indeterminate
	Fluid
	Female
	Male
)

