package modelroller

import (
	"fmt"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/item"
	"github.com/brnsampson/Hissenburg/models/inventory"
    "github.com/charmbracelet/log"
)

func pickFrom[T any](vals []T) T {
	count := len(vals)
	die := dice.New(1, count)
	picked := die.Roll()
	return vals[picked - 1]
}

func RollInv(picker data.ItemBackend, inv *inventory.Inventory) error {
	d20 := dice.New(1, 20)

	armorRoll := d20.Roll()
	armor := item.Empty()
	var err error
	if armorRoll == 20 {
		armor, err = picker.GetItem(item.Armor, "Plate")
		if err != nil {
			return fmt.Errorf("Could not find core game Plate armor")
		}
	} else if armorRoll > 14 {
		armor, err = picker.GetItem(item.Armor, "Chainmail")
		if err != nil {
			return fmt.Errorf("Could not find core game Chainmail armor")
		}
	} else if armorRoll > 4 {
		armor, err = picker.GetItem(item.Armor, "Brigandine")
		if err != nil {
			return fmt.Errorf("Could not find core game Brigandine armor")
		}
	}

	log.Debug("Rolled new armor", "type", armor.Type, "name", armor.Name)

	shieldRoll := d20.Roll()
	helmet := item.Empty()
	shield := item.Empty()
	if shieldRoll == 20 {
		helmet, err = picker.GetItem(item.Armor, "Helmet")
		if err != nil {
			return fmt.Errorf("Could not find core game Helmet")
		}

		shield, err = picker.GetItem(item.Armor, "Shield")
		if err != nil {
			return fmt.Errorf("Could not find core game Shield")
		}
	} else if shieldRoll > 16 {
		shield, err = picker.GetItem(item.Armor, "Shield")
		if err != nil {
			return fmt.Errorf("Could not find core game Shield")
		}
	} else if shieldRoll > 13 {
		helmet, err = picker.GetItem(item.Armor, "Helmet")
		if err != nil {
			return fmt.Errorf("Could not find core game Helmet")
		}
	}

	log.Debug("Rolled new shield", "type", shield.Type, "name", shield.Name)
	log.Debug("Rolled new helmet", "type", helmet.Type, "name", helmet.Name)

	weaponRoll := d20.Roll()
	chosen := ""
	if weaponRoll == 20 {
		chosen = pickFrom([]string{"Halberd", "War Hammer", "Battle Axe"})
	} else if weaponRoll > 14 {
		chosen = pickFrom([]string{"Longbow", "Crossbow", "Sling"})
	} else if weaponRoll > 5 {
		chosen = pickFrom([]string{"Sword", "Mace", "Axe"})
	} else {
		chosen = pickFrom([]string{"Dagger", "Cudgel", "Staff"})
	}

	weapon, err := picker.GetItem(item.Weapon, chosen)
	if err != nil {
		return fmt.Errorf("Could not find core game weapon %s", chosen)
	}
	log.Debug("Rolled new weapon", "type", weapon.Type, "name", weapon.Name)

	gear, err := picker.PickItem(item.Gear)
	if err != nil {
		return err
	}
	log.Debug("Rolled new gear", "type", gear.Type, "name", gear.Name)

	tool, err := picker.PickItem(item.Tool)
	if err != nil {
		return err
	}
	log.Debug("Rolled new tool", "type", tool.Type, "name", tool.Name)

	trinket, err := picker.PickItem(item.Trinket)
	if err != nil {
		return err
	}
	log.Debug("Rolled new trinket", "type", trinket.Type, "name", trinket.Name)

	extraRoll := d20.Roll()
	extraType := pickFrom([]item.ItemType{item.Tool, item.Trinket})
	if extraRoll > 17 {
		extraType = item.SpellBook
	} else if extraRoll > 13 {
		extraType = pickFrom([]item.ItemType{item.Weapon, item.Armor})
	} else if extraRoll > 5 {
		extraType = item.Gear
	}

	extra, err := picker.PickItem(extraType)
	if err != nil {
		return err
	}
	log.Debug("Rolled new extra", "type", extra.Type, "name", extra.Name)

	// Now we actually pack the inventory... fun fact, some combinations of items won't actually fit!
	for _, thing := range []item.Item{extra, gear, tool, trinket} {
		err := inv.AddToBackpack(thing)
		if err != nil {
			return err
		}
	}

	for _, thing := range []item.Item{shield, helmet, armor, weapon} {
		err := inv.AddActive(thing)
		if err != nil {
			return err
		}
	}

	return nil
}

