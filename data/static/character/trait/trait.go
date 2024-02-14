package trait

import o "github.com/brnsampson/optional"

//go:generate stringer -type=Physique -trimprefix Physique
type Physique int

func NoPhysique() o.Option[Physique] {
	return o.None[Physique]()
}

const (
	PhysiqueAthletic Physique = iota
	PhysiqueBrawny
	PhysiqueFlabby
	PhysiqueLanky
	PhysiqueRugged
	PhysiqueScrawny
	PhysiqueShort
	PhysiqueStatuesque
	PhysiqueStout
	PhysiqueTowering
)

func (n Physique) Count() int {
	return 10
}

//go:generate stringer -type=Skin -trimprefix Skin
type Skin int

func NoSkin() o.Option[Skin] {
	return o.None[Skin]()
}

const (
	SkinBirthmark Skin = iota
	SkinDark
	SkinElongated
	SkinPockmarked
	SkinRosy
	SkinRound
	SkinSoft
	SkinTanned
	SkinTattooed
	SkinWeathered
)

func (n Skin) Count() int {
	return 10
}

//go:generate stringer -type=Hair -trimprefix Hair
type Hair int

func NoHair() o.Option[Hair] {
	return o.None[Hair]()
}

const (
	HairBald Hair = iota
	HairBraided
	HairCurly
	HairFilthy
	HairFrizzy
	HairLong
	HairLuxurious
	HairOily
	HairWavy
	HairWispy
)

func (n Hair) Count() int {
	return 10
}

//go:generate stringer -type=Face -trimprefix Face
type Face int

func NoFace() o.Option[Face] {
	return o.None[Face]()
}

const (
	FaceBony Face = iota
	FaceBroken
	FaceChiseled
	FaceElongated
	FacePale
	FacePerfect
	FaceRatlike
	FaceSharp
	FaceSquare
	FaceSunken
)

func (n Face) Count() int {
	return 10
}

//go:generate stringer -type=Speech -trimprefix Speech
type Speech int

func NoSpeech() o.Option[Speech] {
	return o.None[Speech]()
}

const (
	SpeechBlunt Speech = iota
	SpeechBooming
	SpeechCryptic
	SpeechDroning
	SpeechFormal
	SpeechGravelly
	SpeechPrecise
	SpeechSqueaky
	SpeechStuttering
	SpeechWhispery
)

func (n Speech) Count() int {
	return 10
}

//go:generate stringer -type=Clothing -trimprefix Clothing
type Clothing int

func NoClothing() o.Option[Clothing] {
	return o.None[Clothing]()
}

const (
	ClothingAntique Clothing = iota
	ClothingBloody
	ClothingElegant
	ClothingFilthy
	ClothingForeign
	ClothingFrayed
	ClothingFrumpy
	ClothingLivery
	ClothingRancid
	ClothingSoiled
)

func (n Clothing) Count() int {
	return 10
}

//go:generate stringer -type=Virtue -trimprefix Virtue
type Virtue int

func NoVirtue() o.Option[Virtue] {
	return o.None[Virtue]()
}

const (
	VirtueAmbitious Virtue = iota
	VirtueCautious
	VirtueCourageous
	VirtueDiciplined
	VirtueGregarious
	VirtueHonorable
	VirtueHumble
	VirtueMerciful
	VirtueSerene
	VirtueTollerant
)

func (n Virtue) Count() int {
	return 10
}

//go:generate stringer -type=Vice -trimprefix Vice
type Vice int

func NoVice() o.Option[Vice] {
	return o.None[Vice]()
}

const (
	ViceAggressive Vice = iota
	ViceBitter
	ViceCraven
	ViceDeceitful
	ViceGreedy
	ViceLazy
	ViceNervous
	ViceRude
	ViceVain
	ViceVengeful
)

func (n Vice) Count() int {
	return 10
}

//go:generate stringer -type=Reputation -trimprefix Reputation
type Reputation int

func NoReputation() o.Option[Reputation] {
	return o.None[Reputation]()
}

const (
	ReputationAbitious Reputation = iota
	ReputationBoor
	ReputationDangerous
	ReputationEntertainer
	ReputationHonest
	ReputationLoafer
	ReputationOddball
	ReputationRepulsive
	ReputationRespected
	ReputationWise
)

func (n Reputation) Count() int {
	return 10
}

//go:generate stringer -type=Misfortune -trimprefix Misfortune
type Misfortune int

func NoMisfortune() o.Option[Misfortune] {
	return o.None[Misfortune]()
}

const (
	MisfortuneAbandoned Misfortune = iota
	MisfortuneAddicted
	MisfortuneBlackmailed
	MisfortuneCondemed
	MisfortuneCursed
	MisfortuneDefrauded
	MisfortuneDemoted
	MisfortuneDiscredited
	MisfortuneDisowned
	MisfortuneExiled
)

func (n Misfortune) Count() int {
	return 10
}
