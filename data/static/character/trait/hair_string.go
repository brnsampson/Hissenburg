// Code generated by "stringer -type=Hair -trimprefix Hair"; DO NOT EDIT.

package trait

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[HairBald-0]
	_ = x[HairBraided-1]
	_ = x[HairCurly-2]
	_ = x[HairFilthy-3]
	_ = x[HairFrizzy-4]
	_ = x[HairLong-5]
	_ = x[HairLuxurious-6]
	_ = x[HairOily-7]
	_ = x[HairWavy-8]
	_ = x[HairWispy-9]
}

const _Hair_name = "BaldBraidedCurlyFilthyFrizzyLongLuxuriousOilyWavyWispy"

var _Hair_index = [...]uint8{0, 4, 11, 16, 22, 28, 32, 41, 45, 49, 54}

func (i Hair) String() string {
	if i < 0 || i >= Hair(len(_Hair_index)-1) {
		return "Hair(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Hair_name[_Hair_index[i]:_Hair_index[i+1]]
}
