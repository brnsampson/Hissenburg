package modelroller

import (
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/character"
    "github.com/charmbracelet/log"
)

func RollChar(cpicker data.CharBackend, ipicker data.ItemBackend, c *character.Character) error {
	log.Debug("Character roll requested")
	c.Gender = cpicker.PickGender()
	c.Name = cpicker.PickName(c.Gender)
	c.Surname = cpicker.PickSurname()
	c.Background = cpicker.PickBackground()

	ageDice := dice.New(2, 20)
	c.Age = uint16(ageDice.Roll() + 10)
	RollTraits(cpicker, &c.Traits)
	RollStatus(&c.Status)
	err := RollInv(ipicker, &c.Inventory)
	if err != nil {
		return err
	}

	log.Debug("Character roll completed", "character", c)

	return nil
}
