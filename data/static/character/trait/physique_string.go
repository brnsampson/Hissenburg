// Code generated by "stringer -type=Physique -trimprefix Physique"; DO NOT EDIT.

package trait

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PhysiqueAthletic-0]
	_ = x[PhysiqueBrawny-1]
	_ = x[PhysiqueFlabby-2]
	_ = x[PhysiqueLanky-3]
	_ = x[PhysiqueRugged-4]
	_ = x[PhysiqueScrawny-5]
	_ = x[PhysiqueShort-6]
	_ = x[PhysiqueStatuesque-7]
	_ = x[PhysiqueStout-8]
	_ = x[PhysiqueTowering-9]
}

const _Physique_name = "AthleticBrawnyFlabbyLankyRuggedScrawnyShortStatuesqueStoutTowering"

var _Physique_index = [...]uint8{0, 8, 14, 20, 25, 31, 38, 43, 53, 58, 66}

func (i Physique) String() string {
	if i < 0 || i >= Physique(len(_Physique_index)-1) {
		return "Physique(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Physique_name[_Physique_index[i]:_Physique_index[i+1]]
}
