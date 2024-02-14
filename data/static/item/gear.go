package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=Gear -linecomment
type Gear int

const (
	AirBladder Gear = iota // Air Bladder
	Antitoxin
	Cart // Cart
	Chain // Chain (10ft)
	DowsingRod // Dowsing Rod
	FireOil // Fire Oil
	GrapplingHook // Grappling Hook
	LargeSack // Large Sack
	LargeTrap // Large Trap
	Lockpicks
	Manacles
	Pick
	Pole // Pole (10ft)
	Pully
	Repellent
	Rope // Rope (25ft)
	SpiritWard // Spirit Ward
	Spyglass
	Tinderbox
	Wolfsbane
)

func (n Gear) Count() int {
	return 20
}

func GearFromString(name string) (item.Item, error) {
		gear, ok := GearLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := GearList[gear]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var GearLookup map[string]Gear = map[string]Gear{
	AirBladder.String(): AirBladder,
	Antitoxin.String(): Antitoxin,
	Cart.String(): Cart,
	Chain.String(): Chain,
	DowsingRod.String(): DowsingRod,
	FireOil.String(): FireOil,
	GrapplingHook.String(): GrapplingHook,
	LargeSack.String(): LargeSack,
	LargeTrap.String(): LargeTrap,
	Lockpicks.String(): Lockpicks,
	Manacles.String(): Manacles,
	Pick.String(): Pick,
	Pole.String(): Pole,
	Pully.String(): Pully,
	Repellent.String(): Repellent,
	Rope.String(): Rope,
	SpiritWard.String(): SpiritWard,
	Spyglass.String(): Spyglass,
	Tinderbox.String(): Tinderbox,
	Wolfsbane.String(): Wolfsbane,
}

func newGear(gear Gear, description string, value, storage, size int, stacks bool, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: gear.String(),
		Type: item.Gear,
		Description: o.Some(description),
		Value: uint(value),
		Damage: o.None[dice.Dice](),
		Armor: 0,
		Storage: storage,
		Size: size,
		ActiveSize: size,
		ActiveSlot: item.HandSlot,
		Stackable: stacks,
		Icon: icon,
	}
}

var GearList map[Gear]item.Item = map[Gear]item.Item{
	AirBladder: newGear(
		AirBladder,
		"This is basically a professional whoopee cushion. Maybe you can find a creative use for it!",
		5,
		0,
		1,
		false,
		o.None[string](),
	),
	Antitoxin: newGear(
		Antitoxin,
		"Strangely, anti-toxin cancels out poison or venoms.",
		0,
		0,
		1,
		false,
		o.None[string](),
	),
	Cart: newGear(
		Cart,
		"A cart small enough to be pushed by hand (+4 slots, bulky). Look, corporate really needs you to pull your weight",
		30,
		4,
		2,
		false,
		o.None[string](),
	),
	Chain: newGear(
		Chain,
		"A 10ft metal chain",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
	DowsingRod: newGear(
		DowsingRod,
		"A rod used to dowse, which is when an idiot looks for buried water by holding a stick",
		0,
		0,
		1,
		false,
		o.None[string](),
	),
	FireOil: newGear( // Fire Oil
		FireOil,
		"Oil that burns real good",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
	GrapplingHook: newGear( // Grappling Hook
		GrapplingHook,
		"A rope with a hook tied to one end. If you toss it over something it'll get stuck and you can climb! At least, that's how it works in the movies",
		25,
		0,
		1,
		false,
		o.None[string](),
	),
	LargeSack: newGear(
		LargeSack,
		"You have a large sack. Stick something in there!",
		5,
		0,
		1,
		false,
		o.None[string](),
	),
	LargeTrap: newGear(
		LargeTrap,
		"Curious is the trap-maker's art. Never there to see the efficacy of their own design",
		20,
		0,
		1,
		false,
		o.None[string](),
	),
	Lockpicks: newGear(
		Lockpicks,
		"You can pick a lock or pick your nose, but you can't pick your friends!",
		25,
		0,
		1,
		false,
		o.None[string](),
	),
	Manacles: newGear(
		Manacles,
		"Wrist and/or ankle bindings connected with a strong chain. Kinky!",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
	Pick: newGear(
		Pick,
		"What are picks even really used for?",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
	Pole: newGear( // Pole (10ft)
		Pole,
		"A 10ft pole. How does this fit in your backpack?",
		5,
		0,
		1,
		false,
		o.None[string](),
	),
	Pully: newGear(
		Pully,
		"A pully is one of the simple machines described by Archimedes! Wait, what do you mean you don't care?",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
	Repellent: newGear(
		Repellent,
		"I dunno, if you ask me you're already repellent enough! OOOOOOH, SICK BURN BRO",
		0,
		0,
		1,
		false,
		o.None[string](),
	),
	Rope: newGear( // Rope (25ft)
		Rope,
		"25ft rope. Do all the other objects that need rope come with it? Unclear",
		5,
		0,
		1,
		false,
		o.None[string](),
	),
	SpiritWard: newGear(
		SpiritWard,
		"I don't ALWAYS have something witty to write here. Or know what the thing is supposed to be really",
		0,
		0,
		1,
		false,
		o.None[string](),
	),
	Spyglass: newGear(
		Spyglass,
		"An old-timey way to say telescope",
		40,
		0,
		1,
		false,
		o.None[string](),
	),
	Tinderbox: newGear(
		Tinderbox,
		"A small kit you can use to start fires, you little pyro",
		5,
		0,
		1,
		false,
		o.None[string](),
	),
	Wolfsbane: newGear(
		Wolfsbane,
		"Another situation where I am left wondering what this really is",
		10,
		0,
		1,
		false,
		o.None[string](),
	),
}
