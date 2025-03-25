package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=Trinket --linecomment
type Trinket int

const (
	Bottle Trinket = iota
	CardDeck
	DiceSet
	FacePaint // Face Paint
	FakeJewels // Fake Jewels
	Horn
	Incense
	Instrument
	Lens
	Marbles
	Mirror
	Perfume
	QuillAndInk // Quill and Ink
	SaltPack // Salt Pack
	SmallBell // Small Bell
	Soap
	Sponge
	TarPot //Tar Pot
	Twine
	Whistle
)

func (n Trinket) Count() int {
	return 20
}

func TrinketFromString(name string) (item.Item, error) {
		gear, ok := TrinketLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := TrinketList[gear]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var TrinketLookup map[string]Trinket = map[string]Trinket{
	Bottle.String(): Bottle,
	CardDeck.String(): CardDeck,
	DiceSet.String(): DiceSet,
	FacePaint.String(): FacePaint,
	FakeJewels.String(): FakeJewels,
	Horn.String(): Horn,
	Incense.String(): Incense,
	Instrument.String(): Instrument,
	Lens.String(): Lens,
	Marbles.String(): Marbles,
	Mirror.String(): Mirror,
	Perfume.String(): Perfume,
	QuillAndInk.String(): QuillAndInk,
	SaltPack.String(): SaltPack,
	SmallBell.String(): SmallBell,
	Soap.String(): Soap,
	Sponge.String(): Sponge,
	TarPot.String(): TarPot,
	Twine.String(): Twine,
	Whistle.String(): Whistle,
}

func newTrinket(trinket Trinket, description string, value int, stacks bool, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: trinket.String(),
		Type: item.Trinket,
		Description: o.Some(description),
		Value: uint(value),
		Damage: o.None[dice.Dice](),
		Armor: 0,
		Storage: 0,
		Size: 1,
		ActiveSize: 1,
		ActiveSlot: item.Hand,
		Stackable: stacks,
		Icon: icon,
	}
}

var TrinketList map[Trinket]item.Item = map[Trinket]item.Item{
	Bottle: newTrinket(
		Bottle,
		"A bottle with cap or stopper. What are you going to fill it with?!?! I'm so excited!",
		0,
		false,
		o.None[string](),
	),
	CardDeck: newTrinket(
		CardDeck,
		"What kind of cards? Playing Cards? Tarrot Cards? Stolen credit cards?",
		0,
		true,
		o.None[string](),
	),
	DiceSet: newTrinket(
		DiceSet,
		"At least, like, three or four dice.",
		0,
		true,
		o.None[string](),
	),
	FacePaint: newTrinket(
		FacePaint,
		"You could disguise yourself as anyone who also wears face paint!",
		0,
		false,
		o.None[string](),
	),
	FakeJewels: newTrinket(
		FakeJewels,
		"Pretty, but worthless. Just like you, honey.",
		0,
		true,
		o.None[string](),
	),
	Horn: newTrinket(
		Horn,
		"Bugle, trumpet, or french, it doesn't matter. Blow it like you love it.",
		0,
		false,
		o.None[string](),
	),
	Incense: newTrinket(
		Incense,
		"Smells good! Hey bruh, you got a lighter?",
		0,
		true,
		o.None[string](),
	),
	Instrument: newTrinket(
		Instrument,
		"Just don't be that guy at a party that pulls out an acoustic guitar.",
		0,
		false,
		o.None[string](),
	),
	Lens: newTrinket(
		Lens,
		"Concave or convex, or whatever you need at the moment! Lenses are magic, just like magnets.",
		0,
		false,
		o.None[string](),
	),
	Marbles: newTrinket(
		Marbles,
		"The bag has Jelle printed on it. Who is Jelle?",
		0,
		true,
		o.None[string](),
	),
	Mirror: newTrinket(
		Mirror,
		"Mirror, mirror in the hand, who's the baddest in the land?",
		0,
		false,
		o.None[string](),
	),
	Perfume: newTrinket(
		Perfume,
		"PSA: perfume and cologne are not replacements for showers.",
		0,
		false,
		o.None[string](),
	),
	QuillAndInk: newTrinket(
		QuillAndInk,
		"Note from your GM: please take notes.",
		0,
		false,
		o.None[string](),
	),
	SaltPack: newTrinket(
		SaltPack,
		"Also known as a flavor pack.",
		0,
		true,
		o.None[string](),
	),
	SmallBell: newTrinket(
		SmallBell,
		"Ideal for tying around the neck of a small, cute animal.",
		0,
		false,
		o.None[string](),
	),
	Soap: newTrinket(
		Soap,
		"Here you go, you filthy animals.",
		0,
		true,
		o.None[string](),
	),
	Sponge: newTrinket(
		Sponge,
		"Like the kind you use for cleaning, not the kind from the ocean. Wait, THEY ARE THE SAME THINGS?!?!?",
		0,
		false,
		o.None[string](),
	),
	TarPot: newTrinket(
		TarPot,
		"Please don't make a mess.",
		0,
		false,
		o.None[string](),
	),
	Twine: newTrinket(
		Twine,
		"Twine is a twisted pair of fiber. Multiple pieces of twine can be used to make rope. Interesting, right?",
		0,
		true,
		o.None[string](),
	),
	Whistle: newTrinket(
		Whistle,
		"I would ask you to blow the whistle, but you might be too short.",
		0,
		false,
		o.None[string](),
	),
}
