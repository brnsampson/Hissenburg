package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=Armor --linecomment
type Armor int

const (
	Shield Armor = iota
	Helmet
	Gambeson
	Brigandine
	Chainmail
	Plate
)

func (n Armor) Count() int {
	return 6
}

func ArmorFromString(name string) (item.Item, error) {
		gear, ok := ArmorLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := ArmorList[gear]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var ArmorLookup map[string]Armor = map[string]Armor{
	Shield.String(): Shield,
	Helmet.String(): Helmet,
	Gambeson.String(): Gambeson,
	Brigandine.String(): Brigandine,
	Chainmail.String(): Chainmail,
	Plate.String(): Plate,
}

func newArmor(armor Armor, description string, value, armorVal, size int, slot item.ItemSlot, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: armor.String(),
		Type: item.Armor,
		Description: o.Some(description),
		Value: uint(value),
		Damage: o.None[dice.Dice](),
		Armor: armorVal,
		Storage: 0,
		Size: size,
		ActiveSize: size,
		ActiveSlot: slot,
		Stackable: false,
		Icon: icon,
	}
}

var ArmorList map[Armor]item.Item = map[Armor]item.Item{
	Shield: newArmor(
		Shield,
		"A sturdy shield which can be held in one hand and used to (hopefully) deflect attacks.",
		10,
		1,
		1,
		item.Hand,
		o.None[string](),
	),
	Helmet: newArmor(
		Helmet,
		"Head protection. Must be worn on the upper body to be useful.",
		10,
		1,
		1,
		item.Head,
		o.None[string](),
	),
	Gambeson: newArmor(
		Gambeson,
		"A long jacket made of multiple layers of quilted fabric. A basic, but dependable armor.",
		15,
		1,
		1,
		item.Torso,
		o.None[string](),
	),
	Brigandine: newArmor(
		Brigandine,
		"A jacket or doublet with many small, overlapping metal plates rivited to it. The plates may be between layers of cloth or exposed in the front of the armor.",
		20,
		1,
		2,
		item.Torso,
		o.None[string](),
	),
	Chainmail: newArmor(
		Chainmail,
		"A flexible armor made up of many rings rivited together into the shape of a long coat or tunic.",
		40,
		2,
		2,
		item.Torso,
		o.None[string](),
	),
	Plate: newArmor(
		Plate,
		"Armor made of large metal plates shaped to fit the wearer's body.",
		60,
		3,
		2,
		item.Torso,
		o.None[string](),
	),
}
