package trait

//go:generate stringer -type=TraitType
type TraitType int

const (
	Physique TraitType = iota
	Skin
	Hair
	Face
	Speech
	Clothing
	Virtue
	Vice
	Reputation
	Misfortune
)
