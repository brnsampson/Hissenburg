package picker

import (
	"math/rand"
	"github.com/brnsampson/Hissenburg/data/static/character/name"
	"github.com/brnsampson/Hissenburg/data/static/character/gender"
)

type Name interface {
	Id() int
	String() string
}

type Countable interface {
	~int
	Count() int
}

// PickRandom will pick a random value for some type that based on an int and takes the values (0, Count()). If there are
// no special rules or logic that go into picking a particular value.
func PickRandom[T any](from []T) T {
	c := len(from)
	picked := rand.Intn(c)
	return from[picked]
}

//func PickRandom[T Countable]() T {
//	var tmp T
//	c := tmp.Count()
//	picked := rand.Intn(c)
//	return T(picked)
//}

func RandomGender() gender.Gender {
	// Exclude gender.Undefined... That is more meant to indicate no selection has been made
	var g gender.Gender
	c := g.Count()
	picked := rand.Intn(c - 1) + 1
	return gender.Gender(picked)
}

func RandomName(g gender.Gender) Name  {
	if g == gender.Intersex || g == gender.Indeterminate {
		var n name.AsexName
		c := n.Count()
		picked := rand.Intn(c)
		n = name.AsexName(picked)
		return n
	} else if g == gender.Male {
		var n name.MaleName
		c := n.Count()
		picked := rand.Intn(c)
		n = name.MaleName(picked)
		return n
	} else if g == gender.Female {
		var n name.FemaleName
		c := n.Count()
		picked := rand.Intn(c)
		n = name.FemaleName(picked)
		return n
	}

	// The gender.Undefined and gender.Fluid cases can select from any name right now.
	var m name.MaleName
	var f name.FemaleName
	var a name.AsexName
	cm := m.Count()
	cf := f.Count()
	ca := a.Count()
	picked := rand.Intn(cm + cf + ca)
	if tmp := picked - (cm + cf); tmp >= 0 {
		// asexual name picked
		a = name.AsexName(tmp)
		return a
	} else if tmp := picked - cm; tmp >= 0 {
		// female name picked
		f = name.FemaleName(tmp)
		return f
	} else {
		// male name picked
		m = name.MaleName(picked)
		return m
	}
}
