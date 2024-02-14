package modelroller

import (
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/status"
)

func RollStatus(s *status.Status) {
	hpDice := dice.New(1, 6)
	s.MaxHP = uint8(hpDice.Roll())

	strDice := dice.New(3, 6)
	s.MaxStr = uint8(strDice.Roll())

	dexDice := dice.New(3, 6)
	s.MaxDex = uint8(dexDice.Roll())

	willDice := dice.New(3, 6)
	s.MaxWill = uint8(willDice.Roll())

	s.HP = s.MaxHP
	s.Str = s.MaxStr
	s.Dex = s.MaxDex
	s.Will = s.MaxWill
}
