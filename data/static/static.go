package static

import (
	"strings"
	"fmt"
	"math/rand"
	"github.com/brnsampson/Hissenburg/data/static/character/name"
	"github.com/brnsampson/Hissenburg/data/static/character/background"
	"github.com/brnsampson/Hissenburg/data/static/character/trait"
	"github.com/brnsampson/Hissenburg/data/static/item"
	"github.com/brnsampson/Hissenburg/models/character"
	t "github.com/brnsampson/Hissenburg/models/trait"
	i "github.com/brnsampson/Hissenburg/models/item"
)

type countable interface {
	~int
	Count() int
}

type countableListable interface {
	countable
	String() string
}

type StaticBackend struct {}

func New() StaticBackend {
	return StaticBackend{}
}

func pickRandom[T countable]() T {
	var tmp T
	c := tmp.Count()
	picked := rand.Intn(c)
	return T(picked)
}

func list[T countableListable](filter string) []T {
	var t T
	c := t.Count()

	tmp := make([]T, 0)
	for i := 1; i <= c; i ++ {
		tmpstr := T(i).String()
		if strings.HasPrefix(tmpstr, filter) {
			tmp[i] = T(i)
		}
	}
	return tmp
}

func listString[T countableListable](filter string) []string {
	var t T
	c := t.Count()

	tmp := make([]string, 0)
	for i := 1; i <= c; i ++ {
		tmpstr := T(i).String()
		if strings.HasPrefix(tmpstr, filter) {
			tmp[i] = tmpstr
		}
	}
	return tmp
}

func (sb StaticBackend) ListGender(filter string) []character.Gender {
	var g character.Gender
	c := g.Count()
	genders := make([]character.Gender, 0)
	for i := 1; i <= c; i ++ {
		if strings.HasPrefix(character.Gender(i).String(), filter) {
			genders[i] = character.Gender(i)

		}
	}
	return genders
}

func (sb StaticBackend) PickGender() character.Gender {
	// Exclude gender.Undefined... That is more meant to indicate no selection has been made
	var g character.Gender
	c := g.Count()
	picked := rand.Intn(c - 1) + 1
	return character.Gender(picked)
}

func (sb StaticBackend) ListName(gender character.Gender, filter string) []string {
	names := make([]string, 0)

	if gender != character.Female {
		var n name.MaleName
		maleCount := n.Count()
		for i := 0; i <= maleCount; i++ {
			names = append(names, name.MaleName(i).String())
		}
	}

	if gender != character.Male {
		var n name.FemaleName
		femaleCount := n.Count()
		for i := 0; i <= femaleCount; i++ {
			names = append(names, name.FemaleName(i).String())
		}
	}

	var n name.AsexName
	asexCount := n.Count()
	for i := 0; i <= asexCount; i++ {
		names = append(names, name.AsexName(i).String())
	}

	filtered := make([]string, 0)
	for _, n := range names {
		if strings.HasPrefix(n, filter) {
			filtered = append(filtered, n)
		}
	}

	return filtered
}

func (sb StaticBackend) PickName(gender character.Gender) string {
	if gender == character.Male {
		var n name.MaleName
		c := n.Count()
		picked := rand.Intn(c)
		n = name.MaleName(picked)
		return n.String()
	} else if gender == character.Female {
		var n name.FemaleName
		c := n.Count()
		picked := rand.Intn(c)
		n = name.FemaleName(picked)
		return n.String()
	}

	// Any gender outside male or female can select from any name right now
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
		return a.String()
	} else if tmp := picked - cm; tmp >= 0 {
		// female name picked
		f = name.FemaleName(tmp)
		return f.String()
	} else {
		// male name picked
		m = name.MaleName(picked)
		return m.String()
	}
}

func (sb StaticBackend) ListSurname(filter string) []string {
	return listString[name.Surname](filter)
}

func (sb StaticBackend) PickSurname() string {
	return pickRandom[name.Surname]().String()
}

func (sb StaticBackend) ListBackground(filter string) []string {
	return listString[background.Background](filter)
}

func (sb StaticBackend) PickBackground() string {
	return pickRandom[background.Background]().String()
}

func (sb StaticBackend) ListTrait(kind t.TraitType, filter string) []string {
	if kind == t.Physique {
		return listString[trait.Physique](filter)
	} else if kind == t.Skin {
		return listString[trait.Skin](filter)
	} else if kind == t.Hair {
		return listString[trait.Hair](filter)
	} else if kind == t.Face {
		return listString[trait.Face](filter)
	} else if kind == t.Speech {
		return listString[trait.Speech](filter)
	} else if kind == t.Clothing {
		return listString[trait.Clothing](filter)
	} else if kind == t.Virtue {
		return listString[trait.Virtue](filter)
	} else if kind == t.Vice {
		return listString[trait.Vice](filter)
	} else if kind == t.Reputation {
		return listString[trait.Reputation](filter)
	} else if kind == t.Misfortune {
		return listString[trait.Misfortune](filter)
	} else {
		return make([]string, 0)
	}
}

func (sb StaticBackend) PickTrait(kind t.TraitType) string {
	if kind == t.Physique {
		return pickRandom[trait.Physique]().String()
	} else if kind == t.Skin {
		return pickRandom[trait.Skin]().String()
	} else if kind == t.Hair {
		return pickRandom[trait.Hair]().String()
	} else if kind == t.Face {
		return pickRandom[trait.Face]().String()
	} else if kind == t.Speech {
		return pickRandom[trait.Speech]().String()
	} else if kind == t.Clothing {
		return pickRandom[trait.Clothing]().String()
	} else if kind == t.Virtue {
		return pickRandom[trait.Virtue]().String()
	} else if kind == t.Vice {
		return pickRandom[trait.Vice]().String()
	} else if kind == t.Reputation {
		return pickRandom[trait.Reputation]().String()
	} else if kind == t.Misfortune {
		return pickRandom[trait.Misfortune]().String()
	} else {
		return ""
	}
}

func (sb StaticBackend) ListAllItems(filter string) []i.Item {
	items := make([]i.Item, 0)
	tmpGear := list[item.Gear](filter)
	for _, toAdd := range tmpGear {
		items = append(items, item.GearList[toAdd])
	}

	tmpTrink := list[item.Trinket](filter)
	for _, toAdd := range tmpTrink {
		items = append(items, item.TrinketList[toAdd])
	}

	tmpArmor := list[item.Armor](filter)
	for _, toAdd := range tmpArmor {
		items = append(items, item.ArmorList[toAdd])
	}

	tmpWeapon := list[item.Weapon](filter)
	for _, toAdd := range tmpWeapon {
		items = append(items, item.WeaponList[toAdd])
	}

	tmpSpell := list[item.SpellBook](filter)
	for _, toAdd := range tmpSpell {
		items = append(items, item.SpellBookList[toAdd])
	}
	return items
}

func (sb StaticBackend) ListItem(kind i.ItemType, filter string) []i.Item {
	items := make([]i.Item, 0)
	if kind == i.Gear {
		tmp := list[item.Gear](filter)
		for _, toAdd := range tmp {
			items = append(items, item.GearList[toAdd])
		}
	} else if kind == i.Trinket {
		tmp := list[item.Trinket](filter)
		for _, toAdd := range tmp {
			items = append(items, item.TrinketList[toAdd])
		}
	} else if kind == i.Armor {
		tmp := list[item.Armor](filter)
		for _, toAdd := range tmp {
			items = append(items, item.ArmorList[toAdd])
		}
	} else if kind == i.Weapon {
		tmp := list[item.Weapon](filter)
		for _, toAdd := range tmp {
			items = append(items, item.WeaponList[toAdd])
		}
	} else if kind == i.SpellBook {
		tmp := list[item.SpellBook](filter)
		for _, toAdd := range tmp {
			items = append(items, item.SpellBookList[toAdd])
		}
	}
	return items
}

func (sb StaticBackend) GetItem(kind i.ItemType, name string) (i.Item, error) {
	if kind == i.Armor {
		return item.ArmorFromString(name)
	} else if kind == i.Weapon {
		return item.WeaponFromString(name)
	} else if kind == i.Gear {
		return item.GearFromString(name)
	} else if kind == i.Trinket {
		return item.TrinketFromString(name)
	} else if kind == i.Tool {
		return item.ToolFromString(name)
	} else if kind == i.SpellBook {
		return item.SpellBookFromString(name)
	}

	// We shouldn't get here unless there is a new type of item we havent covered
	return i.Empty(), fmt.Errorf("Unknown ItemType used")
}

func (sb StaticBackend) PickItem(kind i.ItemType) (i.Item, error) {
	if kind == i.Armor {
		name := pickRandom[item.Armor]().String()
		return item.ArmorFromString(name)
	} else if kind == i.Weapon {
		name := pickRandom[item.Weapon]().String()
		return item.WeaponFromString(name)
	} else if kind == i.Gear {
		name := pickRandom[item.Gear]().String()
		return item.GearFromString(name)
	} else if kind == i.Trinket {
		name := pickRandom[item.Trinket]().String()
		return item.TrinketFromString(name)
	} else if kind == i.Tool {
		name := pickRandom[item.Tool]().String()
		return item.ToolFromString(name)
	} else if kind == i.SpellBook {
		name := pickRandom[item.SpellBook]().String()
		return item.SpellBookFromString(name)
	}

	// We shouldn't get here unless there is a new type of item we havent covered
	return i.Empty(), fmt.Errorf("Unknown ItemType used")
}

