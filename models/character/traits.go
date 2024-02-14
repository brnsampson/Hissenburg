package character

type Traits struct {
	Physique string
	Skin string
	Hair string
	Face string
	Speech string
	Clothing string
	Virtue string
	Vice string
	Reputation string
	Misfortune string
}

func NewTraits() Traits {
	return Traits {
		Physique: "",
		Skin: "",
		Hair: "",
		Face: "",
		Speech: "",
		Clothing: "",
		Virtue: "",
		Vice: "",
		Reputation: "",
		Misfortune: "",
	}
}
