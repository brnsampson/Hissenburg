// Code generated by "stringer -type=Weapon --linecomment"; DO NOT EDIT.

package item

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Dagger-0]
	_ = x[Cudgel-1]
	_ = x[Sickle-2]
	_ = x[Staff-3]
	_ = x[Spear-4]
	_ = x[Sword-5]
	_ = x[Mace-6]
	_ = x[Flail-7]
	_ = x[Axe-8]
	_ = x[Halberd-9]
	_ = x[WarHammer-10]
	_ = x[LongSword-11]
	_ = x[BattleAxe-12]
	_ = x[Sling-13]
	_ = x[Longbow-14]
	_ = x[Crossbow-15]
}

const _Weapon_name = "DaggerCudgelSickleStaffSpearSwordMaceFlailAxeHalberdWar HammerLong SwordBattle AxeSlingLongbowCrossbow"

var _Weapon_index = [...]uint8{0, 6, 12, 18, 23, 28, 33, 37, 42, 45, 52, 62, 72, 82, 87, 94, 102}

func (i Weapon) String() string {
	if i < 0 || i >= Weapon(len(_Weapon_index)-1) {
		return "Weapon(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Weapon_name[_Weapon_index[i]:_Weapon_index[i+1]]
}