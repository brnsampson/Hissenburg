// Code generated by "stringer -type=Armor --linecomment"; DO NOT EDIT.

package item

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Shield-0]
	_ = x[Helmet-1]
	_ = x[Gambeson-2]
	_ = x[Brigandine-3]
	_ = x[Chainmail-4]
	_ = x[Plate-5]
}

const _Armor_name = "ShieldHelmetGambesonBrigandineChainmailPlate"

var _Armor_index = [...]uint8{0, 6, 12, 20, 30, 39, 44}

func (i Armor) String() string {
	if i < 0 || i >= Armor(len(_Armor_index)-1) {
		return "Armor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Armor_name[_Armor_index[i]:_Armor_index[i+1]]
}
