package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=Weapon --linecomment
type Weapon int

func NoWeapon() o.Option[Weapon] {
	return o.None[Weapon]()
}

const (
	Dagger Weapon = iota
	Cudgel
	Sickle
	Staff
	Spear
	Sword
	Mace
	Flail
	Axe
	Halberd
	WarHammer // War Hammer
	LongSword // Long Sword
	BattleAxe // Battle Axe
	Sling
	Longbow
	Crossbow
)

func (n Weapon) Count() int {
	return 16
}

func WeaponFromString(name string) (item.Item, error) {
		w, ok := WeaponLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := WeaponList[w]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var WeaponLookup map[string]Weapon = map[string]Weapon{
	Dagger.String(): Dagger,
	Cudgel.String(): Cudgel,
	Sickle.String(): Sickle,
	Staff.String(): Staff,
	Spear.String(): Spear,
	Sword.String(): Sword,
	Mace.String(): Mace,
	Flail.String(): Flail,
	Axe.String(): Axe,
	Halberd.String(): Halberd,
	WarHammer.String(): WarHammer,
	LongSword.String(): LongSword,
	BattleAxe.String(): BattleAxe,
	Sling.String(): Sling,
	Longbow.String(): Longbow,
	Crossbow.String(): Crossbow,
}

func newWeapon(weapon Weapon, description string, d dice.Dice, value, size int, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: weapon.String(),
		Type: item.Weapon,
		Description: o.Some(description),
		Value: uint(value),
		Damage: o.Some(d),
		Armor: 0,
		Storage: 0,
		Size: size,
		ActiveSize: size,
		ActiveSlot: item.HandSlot,
		Stackable: false,
		Icon: icon,
	}
}


var WeaponList map[Weapon]item.Item = map[Weapon]item.Item{
	Dagger: newWeapon(Dagger, "A long knife or blade", dice.New(1, 6), 5, 1, o.None[string]()),
	Cudgel: newWeapon(Cudgel, "A short, one-handed club", dice.New(1, 6), 5, 1, o.None[string]()),
	Sickle: newWeapon(Sickle, "A one-handed, curved blade normally used to harvest plants", dice.New(1, 6), 5, 1, o.None[string]()),
	Staff: newWeapon(Staff, "A long wooden pole", dice.New(1, 6), 5, 1, o.None[string]()),
	Spear: newWeapon(Spear, "A long wooden handle with a sharpened point at the end. Light enough to be used one-handed", dice.New(1, 8), 10, 1, o.None[string]()),
	Sword: newWeapon(Sword, "A long, one-handed metal blade", dice.New(1, 8), 10, 1, o.None[string]()),
	Mace: newWeapon(Mace, "A one-handed metal club. Often weighted at the end to increase impact", dice.New(1, 8), 10, 1, o.None[string]()),
	Flail: newWeapon(Flail, "Two lengths of wood joined together with chain or rope", dice.New(1, 8), 10, 1, o.None[string]()),
	Axe: newWeapon(Axe, "A sharp, curved blade at the end of a wooden handle. Traditionally used to fell trees", dice.New(1, 8), 10, 1, o.None[string]()),
	Halberd: newWeapon(Halberd, "A long, two-handed polearm with an axe and spear-point at the tip", dice.New(1, 10), 20, 2, o.None[string]()),
	WarHammer: newWeapon(WarHammer, "A long, two-handed polearm with a hammer mounted at the end. On the back of the hammer is often a spike capable of piercing armor", dice.New(1, 10), 20, 1, o.None[string]()),
	LongSword: newWeapon(LongSword, "A larger sword, intended to be used with two hands", dice.New(1, 10), 20, 2, o.None[string]()),
	BattleAxe: newWeapon(BattleAxe, "A large, two-handed axe specifically designed for combat", dice.New(1, 10), 20, 2, o.None[string]()),
	Sling: newWeapon(Sling, "An ancient projectile weapon which is swung around the head and released to launch rocks or lead balls at suprising speed", dice.New(1, 4), 5, 1, o.None[string]()),
	Longbow: newWeapon(Longbow, "A curved length of wood bent and bound with string to make a spring. It fires arrorws", dice.New(1, 6), 20, 2, o.None[string]()),
	Crossbow: newWeapon(Crossbow, "A mechanical bow designed to use metal instead of wood as a spring. Shoots short, heavy bolts", dice.New(1, 8), 30, 2, o.None[string]()),
}
