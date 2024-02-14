package modelroller

import (
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/models/character"
	"github.com/brnsampson/Hissenburg/models/trait"
)

func RollTraits(picker data.CharBackend, t *character.Traits) {
	t.Physique = picker.PickTrait(trait.Physique)
	t.Skin = picker.PickTrait(trait.Skin)
	t.Hair = picker.PickTrait(trait.Hair)
	t.Face = picker.PickTrait(trait.Face)
	t.Speech = picker.PickTrait(trait.Speech)
	t.Clothing = picker.PickTrait(trait.Clothing)
	t.Virtue = picker.PickTrait(trait.Virtue)
	t.Vice = picker.PickTrait(trait.Vice)
	t.Reputation = picker.PickTrait(trait.Reputation)
	t.Misfortune = picker.PickTrait(trait.Misfortune)
}

