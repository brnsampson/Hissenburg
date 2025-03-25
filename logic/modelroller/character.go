package modelroller

import (
	"context"
	"strings"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/character"
    "github.com/charmbracelet/log"
)

func RollStatus(s character.Status) character.Status {
	hpDice := dice.New(1, 6)
	s.MaxHp = int64(hpDice.Roll())

	strDice := dice.New(3, 6)
	s.MaxStr = int64(strDice.Roll())

	dexDice := dice.New(3, 6)
	s.MaxDex = int64(dexDice.Roll())

	willDice := dice.New(3, 6)
	s.MaxWill = int64(willDice.Roll())

	s.Hp = s.MaxHp
	s.Str = s.MaxStr
	s.Dex = s.MaxDex
	s.Will = s.MaxWill

	return s
}

func RollTraits(ctx context.Context, backend data.CharRepo, traits character.Traits) (character.Traits, error) {
	Physique, err := backend.GetRandomPhysique(ctx)
	if err != nil {
		return traits, err
	}
	traits.Physique = Physique.Physique

	Skin, err := backend.GetRandomSkin(ctx)
	if err != nil {
		return traits, err
	}
	traits.Skin = Skin.Skin

	Hair, err := backend.GetRandomHair(ctx)
	if err != nil {
		return traits, err
	}
	traits.Hair = Hair.Hair

	Face, err := backend.GetRandomFace(ctx)
	if err != nil {
		return traits, err
	}
	traits.Face = Face.Face

	Speech, err := backend.GetRandomSpeech(ctx)
	if err != nil {
		return traits, err
	}
	traits.Speech = Speech.Speech

	Clothing, err := backend.GetRandomClothing(ctx)
	if err != nil {
		return traits, err
	}
	traits.Clothing = Clothing.Clothing

	Virtue, err := backend.GetRandomVirtue(ctx)
	if err != nil {
		return traits, err
	}
	traits.Virtue = Virtue.Virtue

	Vice, err := backend.GetRandomVice(ctx)
	if err != nil {
		return traits, err
	}
	traits.Vice = Vice.Vice

	Reputation, err := backend.GetRandomReputation(ctx)
	if err != nil {
		return traits, err
	}
	traits.Reputation = Reputation.Reputation

	Misfortune, err := backend.GetRandomMisfortune(ctx)
	if err != nil {
		return traits, err
	}
	traits.Misfortune = Misfortune.Misfortune

	return traits, nil
}

func RollIdentity(ctx context.Context, backend data.IdentityRepo, ident character.Identity) (character.Identity, error) {
	ageDice := dice.New(2, 20)
	ident.Age = int64(ageDice.Roll() + 10)

	gender, err := backend.GetRandomGender(ctx)
	if err != nil {
		return ident, err
	}
	ident.Gender = gender.Gender

	background, err := backend.GetRandomBackground(ctx)
	if err != nil {
		return ident, err
	}
	ident.Background = background

	surname, err := backend.GetRandomSurname(ctx)
	if err != nil {
		return ident, err
	}
	ident.Surname = surname.Surname

	if strings.ToLower(gender.Gender) == "male" {
		name, err := backend.GetRandomMasculineName(ctx)
		if err != nil {
			return ident, err
		}
		ident.Name = name.Name
	} else if strings.ToLower(gender.Gender) == "female" {
		name, err := backend.GetRandomFeminineName(ctx)
		if err != nil {
			return ident, err
		}
		ident.Name = name.Name
	} else {
		name, err := backend.GetRandomName(ctx)
		if err != nil {
			return ident, err
		}
		ident.Name = name.Name
	}

	return ident, nil
}

func RollCharacter(ctx context.Context, backend data.CharRepo, c character.Character) (character.Character, error) {
	status := character.Status{CharacterID: c.ID}
	status = RollStatus(status)
	c.Status = status
	log.Debug("Rolled status for new character", "cid", c.ID, "status", status)

	traits := character.Traits{CharacterID: c.ID}
	traits, err := RollTraits(ctx, backend, traits)
	if err != nil {
		log.Error("Failed to roll traits for new character", "cid", c.ID, "traits", traits)
		return c, err
	}
	c.Traits = traits
	log.Debug("Rolled traits for new character", "cid", c.ID, "traits", traits)

	ident := character.Identity{CharacterID: c.ID}
	id, err := RollIdentity(ctx, backend, ident)
	if err != nil {
		log.Error("Failed to roll identity for new character", "cid", c.ID, "identity", id)
		return c, err
	}
	c.Identity = id
	log.Debug("Rolled identity for new character", "cid", c.ID, "identity", id)

	return c, nil
}

//func RollChar(cpicker data.CharBackend, ipicker data.ItemBackend, c *character.Character) error {
//	log.Debug("Character roll requested")
//	c.Gender = cpicker.PickGender()
//	c.Name = cpicker.PickName(c.Gender)
//	c.Surname = cpicker.PickSurname()
//	c.Background = cpicker.PickBackground()
//
//	ageDice := dice.New(2, 20)
//	c.Age = uint16(ageDice.Roll() + 10)
//	RollTraits(cpicker, &c.Traits)
//	RollStatus(&c.Status)
//	err := RollInv(ipicker, &c.Inventory)
//	if err != nil {
//		return err
//	}
//
//	log.Debug("Character roll completed", "character", c)
//
//	return nil
//}
