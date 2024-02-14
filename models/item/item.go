package item

import (
	"strings"
	"github.com/brnsampson/Hissenburg/logic/dice"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=ItemType
type ItemType int

// ItemType is used for certain decitions such as what item to prefer to keep in hands when adding new items
const (
	UnknownItemType ItemType = iota
	EmptySlot
	Bulk
	Gear
	Trinket
	Tool
	Armor
	Weapon
	SpellBook
	Relic
)

func (it ItemType) Count() int {
	return 10
}

func ListItemTypes() []ItemType {
	var tmp ItemType
	kinds := make([]ItemType, 0)
	for i := range tmp.Count() {
		kinds = append(kinds, ItemType(i))
	}
	return kinds
}

func TypeFromString(kind string) ItemType {
	lower := strings.ToLower(kind)
	for i := 0; i <= 9; i++ {
		it := ItemType(i)
		if lower == strings.ToLower(it.String()) {
			return it
		}
	}
	return UnknownItemType
}

// Slot determines which part of the inventory the character must place an item in order to receive a bonus.
// Note that all items can be held, obviously, but holding chainmail doesn't give you armor for example.
//
//go:generate stringer -type=ItemSlot
type ItemSlot int

const (
	AnySlot ItemSlot = iota
	BackpackSlot
	TorsoSlot
	HeadSlot
	HandSlot
)

func Empty() Item {
	return Item{
		Count:       0,
		Name:        "Empty",
		Type:        EmptySlot,
		Description: o.None[string](),
		Value:       0,
		Damage:      dice.NoDice(),
		Armor:       0,
		Storage:     0,
		Size:        1,
		ActiveSize:  1,
		ActiveSlot:  AnySlot,
		Stackable:   true,
		Icon:        o.None[string](),
	}
}

func MakeBulk(i Item) Item {
	i.Type = Bulk
	return i
}

func NoItem() o.Option[Item] {
	return o.None[Item]()
}

type Item struct {
	Count       int
	Name        string
	Type        ItemType
	Description o.Option[string]
	Value       uint
	Damage      o.Option[dice.Dice]
	Armor       int
	Storage     int
	Size        int
	ActiveSize  int
	ActiveSlot  ItemSlot
	Stackable   bool
	Icon        o.Option[string]
}
