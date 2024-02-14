package item

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	o "github.com/brnsampson/optional"
)

//go:generate stringer -type=Tool --linecomment
type Tool int

const (
	Bellows Tool = iota
	Bucket
	Caltrops
	Chalk
	Chisel
	CookPots // Cook Pots
	Crowbar
	Drill
	FishingRod // Fishing Rod
	Glue
	Grease
	Hammer
	Hourglass
	MetalFile // Metal File
	Nails
	Net
	Saw
	Sealant
	Shovel
	Tongs
)

func (n Tool) Count() int {
	return 20
}

func ToolFromString(name string) (item.Item, error) {
		gear, ok := ToolLookup[name]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		res, ok := ToolList[gear]
		if ok != true {
			return item.Empty(), fmt.Errorf("Item not found")
		}
		return res, nil
}

var ToolLookup map[string]Tool = map[string]Tool{
	Bellows.String(): Bellows,
	Bucket.String(): Bucket,
	Caltrops.String(): Caltrops,
	Chalk.String(): Chalk,
	Chisel.String(): Chisel,
	CookPots.String(): CookPots,
	Crowbar.String(): Crowbar,
	Drill.String(): Drill,
	FishingRod.String(): FishingRod,
	Glue.String(): Glue,
	Grease.String(): Grease,
	Hammer.String(): Hammer,
	Hourglass.String(): Hourglass,
	MetalFile.String(): MetalFile,
	Nails.String(): Nails,
	Net.String(): Net,
	Saw.String(): Saw,
	Sealant.String(): Sealant,
	Shovel.String(): Shovel,
	Tongs.String(): Tongs,
}

func newTool(tool Tool, description string, value int, stacks bool, icon o.Option[string]) item.Item {
	return item.Item{
		Count: 1,
		Name: tool.String(),
		Type: item.Tool,
		Description: o.Some(description),
		Value: uint(value),
		Damage: o.None[dice.Dice](),
		Armor: 0,
		Storage: 0,
		Size: 1,
		ActiveSize: 1,
		ActiveSlot: item.HandSlot,
		Stackable: stacks,
		Icon: icon,
	}
}

var ToolList map[Tool]item.Item = map[Tool]item.Item{
	Bellows: newTool(Bellows, "A tool used to forcefully blow air. Just like you!", 10, false, o.None[string]()),
	Bucket: newTool(Bucket, "A small bucket with a handle. Small enough to be filled and carried one-handed. There are no holes in the bucket, dear Dinah", 10, false, o.None[string]()),
	Caltrops: newTool(Caltrops, "A set of small, sharp, jack-like objects meant to stab into the feet of anything that should step on them. Kids still like jacks these days, right?", 10, false, o.None[string]()),
	Chalk: newTool(Chalk, "A stick of chalk, perfect for temporary drawings on all sorts of surfaces. Leave evidence of your explits everywhere you go!", 10, false, o.None[string]()),
	Chisel: newTool(Chisel, "A wide, flat blade with a handle behind it. Can be used to chip away at stone, carve designs into surfaces, etc. Leave perminant evidence of your exploits everywhere you go!", 10, false, o.None[string]()),
	CookPots: newTool(CookPots, "A set of flame-proof pots", 10, false, o.None[string]()),
	Crowbar: newTool(Crowbar, "A strong length of metal, perfect for prying or using as a lever. Developed by crows in anchient times", 10, false, o.None[string]()),
	Drill: newTool(Drill, "A tool used to create holes. You may be asking yourself 'what makes a drill different than an auger then?' No? You didn't ask that?", 10, false, o.None[string]()),
	FishingRod: newTool(FishingRod, "Fishing rod and tackle", 10, false, o.None[string]()),
	Glue: newTool(Glue, "A strong adhesive. Which kind? Check the label", 10, false, o.None[string]()),
	Grease: newTool(Grease, "Thick grease used to lubicate machinery. How does something so thick lubricate? Ask an engineer", 10, false, o.None[string]()),
	Hammer: newTool(Hammer, "A one-handed hammer used for carpentry or construction. You're lucky you have it, because when all you have is a hammer, everything looks like a nail", 10, false, o.None[string]()),
	Hourglass: newTool(Hourglass, "A pair of glass bulbs connected by a narrow throat and filled with sand. The sand takes a specific amout of time to pour from one side to the other. How long? Depends on the hourglass", 10, false, o.None[string]()),
	MetalFile: newTool(MetalFile, "Is it a file made of metal or is it meant to file metal? Could be both!", 10, false, o.None[string]()),
	Nails: newTool(Nails, "Wait... these were not provided with the hammer? And the hammer doesn't come with the nails?", 10, false, o.None[string]()),
	Net: newTool(Net, "This could be any kind of net! Cargo net! Fishing net! Fishnets! Whatever you can imagine!", 10, false, o.None[string]()),
	Saw: newTool(Saw, "A tool used to cut wood or other materials in a slower and less dramatic fashon than an axe.", 10, false, o.None[string]()),
	Sealant: newTool(Sealant, "Perfect for preventing water from penetrating your deck or stoping an evil god from re-entering our world for 100 years", 10, false, o.None[string]()),
	Shovel: newTool(Shovel, "A tool for digging through the earth. Can you dig it?", 10, false, o.None[string]()),
	Tongs: newTool(Tongs, "Hold things without using your fingers. FINALLY!", 10, false, o.None[string]()),
}
