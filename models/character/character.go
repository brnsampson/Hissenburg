package character

import (
	"strings"
	"github.com/brnsampson/Hissenburg/models/status"
	"github.com/brnsampson/Hissenburg/models/inventory"
)

//go:generate stringer -type=Gender
type Gender int

func (g Gender) Count() int {
	return 6
}

const (
	GenderUndefined Gender = iota
	Intersex
	Indeterminate
	Fluid
	Female
	Male
)

func GenderFromString(g string) Gender {
	g = strings.ToLower(g)
	if g == "genderundefined" {
		return GenderUndefined
	} else if g == "indeterminate" {
		return Indeterminate
	} else if g == "fluid" {
		return Fluid
	} else if g == "female" {
		return Female
	} else if g == "male" {
		return Male
	} else {
		return GenderUndefined
	}
}

type Character struct {
	Name string
	Surname string
	Gender Gender
	Background string
	Age uint16
	Traits Traits
	Status status.Status
	Description string
	Inventory inventory.Inventory
}

func New() Character {
	return Character{
		Name: "",
		Surname: "",
		Gender: GenderUndefined,
		Background: "",
		Age: 0,
		Traits: NewTraits(),
		Status: status.New(),
		Description: "",
		Inventory: inventory.New(),
	}
}
