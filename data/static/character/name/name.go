package name

//go:generate stringer -type=FemaleName -trimprefix Female
type FemaleName int

func (n FemaleName) Id() int {
	return int(n)
}

func (n FemaleName) Count() int {
	return 19
}

const (
	FemaleAgune FemaleName = iota
	FemaleBeatrice
	FemaleBreagan
	FemaleBronwyn
	FemaleCannora
	FemaleDrelil
	FemaleElgile
	FemaleEsme
	FemaleGroua
	FemaleHenaine
	FemaleLiranne
	FemaleLirathil
	FemaleLisabeth
	FemaleMoralil
	FemaleMorgwin
	FemaleSybil
	FemaleTheune
	FemaleYgwal
	FemaleYslen
)

//go:generate stringer -type=MaleName -trimprefix Male
type MaleName int

func (n MaleName) Id() int {
	return int(n)
}

func (n MaleName) Count() int {
	return 19
}

const (
	MaleArwel MaleName = iota
	MaleBevan
	MaleBoroth
	MaleBorrid
	MaleBreagle
	MaleBreglor
	MaleCanhoreal
	MaleEmrys
	MaleEthex
	MaleGringle
	MaleGrinwit
	MaleGruwid
	MaleGruwth
	MaleGwestin
	MaleMannog
	MaleMelnax
	MaleOrthax
	MaleTriunein
	MaleYirmeor
)

//go:generate stringer -type=AsexName
type AsexName int

func (n AsexName) Id() int {
	return int(n)
}

const (
	Wenlan AsexName = iota
)

func (n AsexName) Count() int {
	return 1
}

//go:generate stringer -type=Surname -trimprefix Surname
type Surname int

func (n Surname) Id() int {
	return int(n)
}

const (
	SurnameAbernathy Surname = iota
	SurnameAddercap
	SurnameBurl
	SurnameCandlewick
	SurnameCormick
	SurnameCrumwaller
	SurnameDunswallow
	SurnameGetri
	SurnameGlass
	SurnameHarkness
	SurnameHarper
	SurnameLoomer
	SurnameMalksmilk
	SurnameSmythe
	SurnameSunderman
	SurnameSwinney
	SurnameThatcher
	SurnameTolmen
	SurnameWeaver
	SurnameWolder
)

func (n Surname) Count() int {
	return 20
}

