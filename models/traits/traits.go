package traits

type Traits struct {
	ID int64
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

func New() Traits {
	return Traits{ ID: -1 }
}
