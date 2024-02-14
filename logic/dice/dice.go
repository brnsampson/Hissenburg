package dice

import (
	"math/rand"
	o "github.com/brnsampson/optional"
)

func New(count, sides int) Dice {
	return Dice{count, sides}
}

func SomeDice(count, sides int) o.Option[Dice] {
	return o.Some(New(count, sides))
}

func NoDice() o.Option[Dice] {
	return o.None[Dice]()
}

type Dice struct {
	Count int
	Sides int
}

func (d Dice) Roll() int {
	total := 0
	for i :=  1; i <= d.Count; i++ {
		total += rand.Intn(d.Sides) + 1
	}
	return total
}
