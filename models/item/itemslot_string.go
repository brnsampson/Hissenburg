// Code generated by "stringer -type=ItemSlot"; DO NOT EDIT.

package item

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnySlot-0]
	_ = x[BackpackSlot-1]
	_ = x[TorsoSlot-2]
	_ = x[HeadSlot-3]
	_ = x[HandSlot-4]
}

const _ItemSlot_name = "AnySlotBackpackSlotTorsoSlotHeadSlotHandSlot"

var _ItemSlot_index = [...]uint8{0, 7, 19, 28, 36, 44}

func (i ItemSlot) String() string {
	if i < 0 || i >= ItemSlot(len(_ItemSlot_index)-1) {
		return "ItemSlot(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemSlot_name[_ItemSlot_index[i]:_ItemSlot_index[i+1]]
}