package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=SpellBook --linecomment
type SpellBook int

const (
	Adhere SpellBook = iota
	Anchor
	AnimateObject // Animate Object
	Anthropomorphize
	ArcaneEye // Arcane Eye
	AstralPrison // Astral Prison
	Attract
	AuditoryIllusion // Auditory Illusion
	Babble
	BaitFlower // Bait Flower
	BeastForm // Beast Form
	Befuddle
	BodySwap // Body Swap
	Charm
	Command
	Comprehend
	ConeOfFoam // Cone of Foam
	ControlPlants // Control Plants
	ControlWeather // Control Weather
	CureWounds // Cure Wounds
	Deafen
	DetectMagic // Detect Magic
	Disassemble
	Disguise
	Displace
	Earthquake
	Elasticity
	ElementalWall // Elemental Wall
	Filch
	Flare
	FogCloud // Fog Cloud
	Frenzy
	Gate
	GravityShift // Gravity Shift
	Greed
	Haste
	Hatred
	HearWhispers // Hear Whispers
	Hover
	Hypnotize
	IcyTouch // Icy Touch
	IdentifyOwner // Identify Owner
	Illuminate
	InvisibleTeather // Invisible Tether
	Knock
	Leap
	LiquidAir // Liquid Air
	MagicDampener // Magic Dampener
	Manse
	MarbleCraze // Marble Craze
	Masquerade
	Miniaturize
	MirrorImage // Mirror Image
	Mirrorwalk
	Multiarm
	NightSphere // Night Sphere
	Objectify
	OozeForm // Ooze Form
	Pacify
	Phobia
	Pit
	PrimalSurge // Primal Surge
	PushPull // Push/Pull
	RaiseDead // Raise Dead
	RaiseSpirit // Raise Spirit
	ReadMind // Read Mind
	Repel
	Scry
	SculptElements // Sculpt Elements
	Sense
	MissileShield // Missile Shield
	Shroud
	Shuffle
	Sleep
	Slick
	SmokeForm // Smoke Form
	Sniff
	Snuff
	Sort
	Spectacle
	Spellsaw
	SpiderClimb // Spider Climb
	SummonCube // Summon Cube
	Swarm
	Telekinesis
	Telepathy
	Teleport
	TargetLure // Target Lure
	Thicket
	SummonIdol // Summon Idol
	TimeControl // Time Control
	TrueSight // True Sight
	Upwell
	Vision
	VisualIllusion // Visual Illusion
	Ward
	Web
	Widget
	WizardMark // Wizard Mark
	XRayVision // X-Ray Vision
)

func (n SpellBook) Count() int {
	return 100
}

func SpellBookFromString(name string) (item.Item, error) {
		gear, ok := SpellBookLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := SpellBookList[gear]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var SpellBookLookup map[string]SpellBook = map[string]SpellBook{
	Adhere.String(): Adhere,
	Anchor.String(): Anchor,
	AnimateObject.String(): AnimateObject,
	Anthropomorphize.String(): Anthropomorphize,
	ArcaneEye.String(): ArcaneEye,
	AstralPrison.String(): AstralPrison,
	Attract.String(): Attract,
	AuditoryIllusion.String(): AuditoryIllusion,
	Babble.String(): Babble,
	BaitFlower.String(): BaitFlower,
	BeastForm.String(): BeastForm,
	Befuddle.String(): Befuddle,
	BodySwap.String(): BodySwap,
	Charm.String(): Charm,
	Command.String(): Command,
	Comprehend.String(): Comprehend,
	ConeOfFoam.String(): ConeOfFoam,
	ControlPlants.String(): ControlPlants,
	ControlWeather.String(): ControlWeather,
	CureWounds.String(): CureWounds,
	Deafen.String(): Deafen,
	DetectMagic.String(): DetectMagic,
	Disassemble.String(): Disassemble,
	Disguise.String(): Disguise,
	Displace.String(): Displace,
	Earthquake.String(): Earthquake,
	Elasticity.String(): Elasticity,
	ElementalWall.String(): ElementalWall,
	Filch.String(): Filch,
	Flare.String(): Flare,
	FogCloud.String(): FogCloud,
	Frenzy.String(): Frenzy,
	Gate.String(): Gate,
	GravityShift.String(): GravityShift,
	Greed.String(): Greed,
	Haste.String(): Haste,
	Hatred.String(): Hatred,
	HearWhispers.String(): HearWhispers,
	Hover.String(): Hover,
	Hypnotize.String(): Hypnotize,
	IcyTouch.String(): IcyTouch,
	IdentifyOwner.String(): IdentifyOwner,
	Illuminate.String(): Illuminate,
	InvisibleTeather.String(): InvisibleTeather,
	Knock.String(): Knock,
	Leap.String(): Leap,
	LiquidAir.String(): LiquidAir,
	MagicDampener.String(): MagicDampener,
	Manse.String(): Manse,
	MarbleCraze.String(): MarbleCraze,
	Masquerade.String(): Masquerade,
	Miniaturize.String(): Miniaturize,
	MirrorImage.String(): MirrorImage,
	Mirrorwalk.String(): Mirrorwalk,
	Multiarm.String(): Multiarm,
	NightSphere.String(): NightSphere,
	Objectify.String(): Objectify,
	OozeForm.String(): OozeForm,
	Pacify.String(): Pacify,
	Phobia.String(): Phobia,
	Pit.String(): Pit,
	PrimalSurge.String(): PrimalSurge,
	PushPull.String(): PushPull,
	RaiseDead.String(): RaiseDead,
	RaiseSpirit.String(): RaiseSpirit,
	ReadMind.String(): ReadMind,
	Repel.String(): Repel,
	Scry.String(): Scry,
	SculptElements.String(): SculptElements,
	Sense.String(): Sense,
	MissileShield.String(): MissileShield,
	Shroud.String(): Shroud,
	Shuffle.String(): Shuffle,
	Sleep.String(): Sleep,
	Slick.String(): Slick,
	SmokeForm.String(): SmokeForm,
	Sniff.String(): Sniff,
	Snuff.String(): Snuff,
	Sort.String(): Sort,
	Spectacle.String(): Spectacle,
	Spellsaw.String(): Spellsaw,
	SpiderClimb.String(): SpiderClimb,
	SummonCube.String(): SummonCube,
	Swarm.String(): Swarm,
	Telekinesis.String(): Telekinesis,
	Telepathy.String(): Telepathy,
	Teleport.String(): Teleport,
	TargetLure.String(): TargetLure,
	Thicket.String(): Thicket,
	SummonIdol.String(): SummonIdol,
	TimeControl.String(): TimeControl,
	TrueSight.String(): TrueSight,
	Upwell.String(): Upwell,
	Vision.String(): Vision,
	VisualIllusion.String(): VisualIllusion,
	Ward.String(): Ward,
	Web.String(): Web,
	Widget.String(): Widget,
	WizardMark.String(): WizardMark,
	XRayVision.String(): XRayVision,
}

func newSpellBook(book SpellBook, description string, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: book.String(),
		Type: item.SpellBook,
		Description: o.Some(description),
		Value: 0,
		Damage: o.None[dice.Dice](),
		Armor: 0,
		Storage: 0,
		Size: 1,
		ActiveSize: 2,
		ActiveSlot: item.HandSlot,
		Stackable: false,
		Icon: icon,
	}
}

var SpellBookList map[SpellBook]item.Item = map[SpellBook]item.Item{
	Adhere: newSpellBook(Adhere, "An object is covered in extremely sticky slime", o.None[string]()),
	Anchor: newSpellBook(Anchor, "A strong wire sprouts from your arms, affixing itself to two points within 50ft on each side", o.None[string]()),
	AnimateObject: newSpellBook(AnimateObject, "An object obeys your commands as best it can", o.None[string]()),
	Anthropomorphize: newSpellBook(Anthropomorphize, "", o.None[string]()),
	ArcaneEye: newSpellBook(ArcaneEye, "", o.None[string]()),
	AstralPrison: newSpellBook(AstralPrison, "", o.None[string]()),
	Attract: newSpellBook(Attract, "", o.None[string]()),
	AuditoryIllusion: newSpellBook(AuditoryIllusion, "", o.None[string]()),
	Babble: newSpellBook(Babble, "", o.None[string]()),
	BaitFlower: newSpellBook(BaitFlower, "", o.None[string]()),
	BeastForm: newSpellBook(BeastForm, "", o.None[string]()),
	Befuddle: newSpellBook(Befuddle, "", o.None[string]()),
	BodySwap: newSpellBook(BodySwap, "", o.None[string]()),
	Charm: newSpellBook(Charm, "", o.None[string]()),
	Command: newSpellBook(Command, "", o.None[string]()),
	Comprehend: newSpellBook(Comprehend, "", o.None[string]()),
	ConeOfFoam: newSpellBook(ConeOfFoam, "", o.None[string]()),
	ControlPlants: newSpellBook(ControlPlants, "", o.None[string]()),
	ControlWeather: newSpellBook(ControlWeather, "", o.None[string]()),
	CureWounds: newSpellBook(CureWounds, "", o.None[string]()),
	Deafen: newSpellBook(Deafen, "", o.None[string]()),
	DetectMagic: newSpellBook(DetectMagic, "", o.None[string]()),
	Disassemble: newSpellBook(Disassemble, "", o.None[string]()),
	Disguise: newSpellBook(Disguise, "", o.None[string]()),
	Displace: newSpellBook(Displace, "", o.None[string]()),
	Earthquake: newSpellBook(Earthquake, "", o.None[string]()),
	Elasticity: newSpellBook(Elasticity, "", o.None[string]()),
	ElementalWall: newSpellBook(ElementalWall, "", o.None[string]()),
	Filch: newSpellBook(Filch, "", o.None[string]()),
	Flare: newSpellBook(Flare, "", o.None[string]()),
	FogCloud: newSpellBook(FogCloud, "", o.None[string]()),
	Frenzy: newSpellBook(Frenzy, "", o.None[string]()),
	Gate: newSpellBook(Gate, "", o.None[string]()),
	GravityShift: newSpellBook(GravityShift, "", o.None[string]()),
	Greed: newSpellBook(Greed, "", o.None[string]()),
	Haste: newSpellBook(Haste, "", o.None[string]()),
	Hatred: newSpellBook(Hatred, "", o.None[string]()),
	HearWhispers: newSpellBook(HearWhispers, "", o.None[string]()),
	Hover: newSpellBook(Hover, "", o.None[string]()),
	Hypnotize: newSpellBook(Hypnotize, "", o.None[string]()),
	IcyTouch: newSpellBook(IcyTouch, "", o.None[string]()),
	IdentifyOwner: newSpellBook(IdentifyOwner, "", o.None[string]()),
	Illuminate: newSpellBook(Illuminate, "", o.None[string]()),
	InvisibleTeather: newSpellBook(InvisibleTeather, "", o.None[string]()),
	Knock: newSpellBook(Knock, "", o.None[string]()),
	Leap: newSpellBook(Leap, "", o.None[string]()),
	LiquidAir: newSpellBook(LiquidAir, "", o.None[string]()),
	MagicDampener: newSpellBook(MagicDampener, "", o.None[string]()),
	Manse: newSpellBook(Manse, "", o.None[string]()),
	MarbleCraze: newSpellBook(MarbleCraze, "", o.None[string]()),
	Masquerade: newSpellBook(Masquerade, "", o.None[string]()),
	Miniaturize: newSpellBook(Miniaturize, "", o.None[string]()),
	MirrorImage: newSpellBook(MirrorImage, "", o.None[string]()),
	Mirrorwalk: newSpellBook(Mirrorwalk, "", o.None[string]()),
	Multiarm: newSpellBook(Multiarm, "", o.None[string]()),
	NightSphere: newSpellBook(NightSphere, "", o.None[string]()),
	Objectify: newSpellBook(Objectify, "", o.None[string]()),
	OozeForm: newSpellBook(OozeForm, "", o.None[string]()),
	Pacify: newSpellBook(Pacify, "", o.None[string]()),
	Phobia: newSpellBook(Phobia, "", o.None[string]()),
	Pit: newSpellBook(Pit, "", o.None[string]()),
	PrimalSurge: newSpellBook(PrimalSurge, "", o.None[string]()),
	PushPull: newSpellBook(PushPull, "", o.None[string]()),
	RaiseDead: newSpellBook(RaiseDead, "", o.None[string]()),
	RaiseSpirit: newSpellBook(RaiseSpirit, "", o.None[string]()),
	ReadMind: newSpellBook(ReadMind, "", o.None[string]()),
	Repel: newSpellBook(Repel, "", o.None[string]()),
	Scry: newSpellBook(Scry, "", o.None[string]()),
	SculptElements: newSpellBook(SculptElements, "", o.None[string]()),
	Sense: newSpellBook(Sense, "", o.None[string]()),
	MissileShield: newSpellBook(MissileShield, "", o.None[string]()),
	Shroud: newSpellBook(Shroud, "", o.None[string]()),
	Shuffle: newSpellBook(Shuffle, "", o.None[string]()),
	Sleep: newSpellBook(Sleep, "", o.None[string]()),
	Slick: newSpellBook(Slick, "", o.None[string]()),
	SmokeForm: newSpellBook(SmokeForm, "", o.None[string]()),
	Sniff: newSpellBook(Sniff, "", o.None[string]()),
	Snuff: newSpellBook(Snuff, "", o.None[string]()),
	Sort: newSpellBook(Sort, "", o.None[string]()),
	Spectacle: newSpellBook(Spectacle, "", o.None[string]()),
	Spellsaw: newSpellBook(Spellsaw, "", o.None[string]()),
	SpiderClimb: newSpellBook(SpiderClimb, "", o.None[string]()),
	SummonCube: newSpellBook(SummonCube, "", o.None[string]()),
	Swarm: newSpellBook(Swarm, "", o.None[string]()),
	Telekinesis: newSpellBook(Telekinesis, "", o.None[string]()),
	Telepathy: newSpellBook(Telepathy, "", o.None[string]()),
	Teleport: newSpellBook(Teleport, "", o.None[string]()),
	TargetLure: newSpellBook(TargetLure, "", o.None[string]()),
	Thicket: newSpellBook(Thicket, "", o.None[string]()),
	SummonIdol: newSpellBook(SummonIdol, "", o.None[string]()),
	TimeControl: newSpellBook(TimeControl, "", o.None[string]()),
	TrueSight: newSpellBook(TrueSight, "", o.None[string]()),
	Upwell: newSpellBook(Upwell, "", o.None[string]()),
	Vision: newSpellBook(Vision, "", o.None[string]()),
	VisualIllusion: newSpellBook(VisualIllusion, "", o.None[string]()),
	Ward: newSpellBook(Ward, "", o.None[string]()),
	Web: newSpellBook(Web, "", o.None[string]()),
	Widget: newSpellBook(Widget, "", o.None[string]()),
	WizardMark: newSpellBook(WizardMark, "", o.None[string]()),
	XRayVision: newSpellBook(XRayVision, "", o.None[string]()),
}
