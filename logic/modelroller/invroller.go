package modelroller

import (
	"context"
	"github.com/brnsampson/Hissenburg/data"
	"github.com/brnsampson/Hissenburg/logic/dice"
	"github.com/brnsampson/Hissenburg/models/inventory"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
    "github.com/charmbracelet/log"
)

func pickFrom[T any](vals []T) T {
	count := len(vals)
	die := dice.New(1, count)
	picked := die.Roll()
	return vals[picked - 1]
}

func RollInventory(ctx context.Context, backend data.InventoryRepo) (inventory.Inventory, error) {
	//empty, err := backend.GetItemFromKindAndName(ctx, "empty", "Empty")
	//if err != nil {
	//	log.Error("Error while fetching empty item from DB")
	//	return inventory.New(), err
	//}
	inv := inventory.New()

	d20 := dice.New(1, 20)
	d8 := dice.New(1, 8)

	armorRoll := d20.Roll()
	var armor *sqlc.ItemView
	if armorRoll == 20 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Plate")
		armor = &tmp
		if err != nil {
			log.Error("Error while fetching Plate Armor from DB")
			return inv, err
		}
	} else if armorRoll > 14 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Chainmail")
		armor = &tmp
		if err != nil {
			log.Error("Error while fetching Chainmail Armor from DB")
			return inv, err
		}
	} else if armorRoll > 4 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Brigandine")
		armor = &tmp
		if err != nil {
			log.Error("Error while fetching Brigandine Armor from DB")
			return inv, err
		}
	}
	if armor != nil {
		log.Debug("Rolled new armor", "kind", armor.Kind, "name", armor.Name)
		inv.Equipment.Torso = armor
	} else {
		log.Debug("Unlucky! No armor rolled.")
	}

	shieldRoll := d20.Roll()
	var helmet *sqlc.ItemView
	var shield *sqlc.ItemView
	if shieldRoll == 20 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Helmet")
		helmet = &tmp
		if err != nil {
			log.Error("Error while fetching Helmet Armor from DB")
			return inv, err
		}

		tmp, err = backend.GetItemFromKindAndName(ctx, "armor", "Shield")
		shield = &tmp
		if err != nil {
			log.Error("Error while fetching Shield Armor from DB")
			return inv, err
		}
	} else if shieldRoll > 16 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Shield")
		shield = &tmp
		if err != nil {
			log.Error("Error while fetching Shield Armor from DB")
			return inv, err
		}
	} else if shieldRoll > 13 {
		tmp, err := backend.GetItemFromKindAndName(ctx, "armor", "Helmet")
		helmet = &tmp
		if err != nil {
			log.Error("Error while fetching Helmet Armor from DB")
			return inv, err
		}
	}

	if shield != nil {
		log.Debug("Rolled new shield", "kind", shield.Kind, "name", shield.Name)
	} else {
		log.Debug("Unlucky! No shield rolled.")
	}

	if helmet != nil {
		log.Debug("Rolled new helmet", "kind", helmet.Kind, "name", helmet.Name)
		if inv.Equipment.Torso != nil && inv.Equipment.Torso.ActiveSize <= 1 {
			inv.Equipment.Head = helmet
		} else {
			err := inv.AddToBackpack(*helmet, inventory.MAX_BACKPACK)
			if err != nil {
				if err.IsOutOfSpace() {
					log.Debug("Item could not fit in inventory and was dropped on ground", "kind", helmet.Kind, "item", helmet.Name)
				} else {
					log.Error("Error while adding item to backpack", "kind", helmet.Kind, "item", helmet.Name)
					return inv, err
				}
			}
		}
	} else {
		log.Debug("Unlucky! No helmet rolled.")
	}

	// Save shield until we find our weapon so we can add things to hands at the same time
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

	weapon, err := backend.GetItemFromKindAndName(ctx, "weapon", chosen)
	if err != nil {
		log.Error("Error while fetching Weapon from DB", "name", chosen)
		return inv, err
	}
	log.Debug("Rolled new weapon", "kind", weapon.Kind, "name", weapon.Name)

	if weapon.ActiveSize == 2 {
		inv.Equipment.LeftHand = &weapon
		if shield != nil {
			err := inv.AddToBackpack(*shield, inventory.MAX_BACKPACK)
			if err != nil {
				if err.IsOutOfSpace() {
					log.Debug("Item could not fit in inventory and was dropped on ground", "kind", shield.Kind, "item", shield.Name)
				} else {
					log.Error("Error while dropping item on the ground", "kind", shield.Kind, "item", shield.Name)
					return inv, err
				}
			}
		}
	} else {
		leftyRoll := d8.Roll()
		if leftyRoll == 1 {
			inv.Equipment.LeftHand = &weapon
			inv.Equipment.RightHand = shield
		} else {
			inv.Equipment.RightHand = &weapon
			inv.Equipment.LeftHand = shield
		}
	}

	tool, err := backend.GetRandomItemFromKind(ctx, "Tool")
	if err != nil {
		log.Error("Error while fetching random Tool from DB")
		return inv, err
	}
	log.Debug("Rolled new tool", "kind", tool.Kind, "name", tool.Name)
	ierr := inv.AddToBackpack(tool, inventory.MAX_BACKPACK)
	if ierr != nil {
		if ierr.IsOutOfSpace() {
			log.Debug("Item could not fit in inventory and was dropped on ground", "kind", tool.Kind, "item", tool.Name)
		} else {
			log.Error("Error while dropping item on the ground", "kind", tool.Kind, "item", tool.Name)
			return inv, err
		}
	}

	trinket, err := backend.GetRandomItemFromKind(ctx, "Trinket")
	if err != nil {
		return inv, err
	}
	log.Debug("Rolled new trinket", "type", trinket.Kind, "name", trinket.Name)

	ierr = inv.AddToBackpack(trinket, inventory.MAX_BACKPACK)
	if ierr != nil {
		if ierr.IsOutOfSpace() {
			log.Debug("Item could not fit in inventory and was dropped on ground", "kind", trinket.Kind, "item", trinket.Name)
		} else {
			log.Error("Error while dropping item on the ground", "kind", trinket.Kind, "item", trinket.Name)
			return inv, err
		}
	}

	extraRoll := d20.Roll()
	extraKind := pickFrom([]string{"Tool", "Trinket"})
	if extraRoll > 17 {
		extraKind = "SpellBook"
	} else if extraRoll > 13 {
		extraKind = pickFrom([]string{"Weapon", "Armor"})
	} else if extraRoll > 5 {
		extraKind = "Gear"
	}

	extra, err := backend.GetRandomItemFromKind(ctx, extraKind)
	if err != nil {
		log.Error("Error while fetching random extra from DB", "kind", extraKind)
		return inv, err
	}
	log.Debug("Rolled new bonus item", "type", extra.Kind, "name", extra.Name)

	ierr = inv.AddToBackpack(extra, inventory.MAX_BACKPACK)
	if ierr != nil {
		if ierr.IsOutOfSpace() {
			log.Debug("Item could not fit in inventory and was dropped on ground", "kind", extra.Kind, "item", extra.Name)
		} else {
			log.Error("Error while dropping item on the ground", "kind", extra.Kind, "item", extra.Name)
			return inv, err
		}
	}

	inv.DropOverweight()
	log.Debug("dropped overweight items after rolling inventory")
	inv.SetCalculatedValues()
	log.Debug("set calculated values after rolling inventory")
	return inv, nil
}

//
//func RollInv(picker data.ItemBackend, inv *inventory.Inventory) error {
//	d20 := dice.New(1, 20)
//
//	armorRoll := d20.Roll()
//	armor := item.Empty()
//	var err error
//	if armorRoll == 20 {
//		armor, err = picker.GetItem(item.Armor, "Plate")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Plate armor")
//		}
//	} else if armorRoll > 14 {
//		armor, err = picker.GetItem(item.Armor, "Chainmail")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Chainmail armor")
//		}
//	} else if armorRoll > 4 {
//		armor, err = picker.GetItem(item.Armor, "Brigandine")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Brigandine armor")
//		}
//	}
//
//	log.Debug("Rolled new armor", "type", armor.Type, "name", armor.Name)
//
//	shieldRoll := d20.Roll()
//	helmet := item.Empty()
//	shield := item.Empty()
//	if shieldRoll == 20 {
//		helmet, err = picker.GetItem(item.Armor, "Helmet")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Helmet")
//		}
//
//		shield, err = picker.GetItem(item.Armor, "Shield")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Shield")
//		}
//	} else if shieldRoll > 16 {
//		shield, err = picker.GetItem(item.Armor, "Shield")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Shield")
//		}
//	} else if shieldRoll > 13 {
//		helmet, err = picker.GetItem(item.Armor, "Helmet")
//		if err != nil {
//			return fmt.Errorf("Could not find core game Helmet")
//		}
//	}
//
//	log.Debug("Rolled new shield", "type", shield.Type, "name", shield.Name)
//	log.Debug("Rolled new helmet", "type", helmet.Type, "name", helmet.Name)
//
//	weaponRoll := d20.Roll()
//	chosen := ""
//	if weaponRoll == 20 {
//		chosen = pickFrom([]string{"Halberd", "War Hammer", "Battle Axe"})
//	} else if weaponRoll > 14 {
//		chosen = pickFrom([]string{"Longbow", "Crossbow", "Sling"})
//	} else if weaponRoll > 5 {
//		chosen = pickFrom([]string{"Sword", "Mace", "Axe"})
//	} else {
//		chosen = pickFrom([]string{"Dagger", "Cudgel", "Staff"})
//	}
//
//	weapon, err := picker.GetItem(item.Weapon, chosen)
//	if err != nil {
//		return fmt.Errorf("Could not find core game weapon %s", chosen)
//	}
//	log.Debug("Rolled new weapon", "type", weapon.Type, "name", weapon.Name)
//
//	gear, err := picker.PickItem(item.Gear)
//	if err != nil {
//		return err
//	}
//	log.Debug("Rolled new gear", "type", gear.Type, "name", gear.Name)
//
//	tool, err := picker.PickItem(item.Tool)
//	if err != nil {
//		return err
//	}
//	log.Debug("Rolled new tool", "type", tool.Type, "name", tool.Name)
//
//	trinket, err := picker.PickItem(item.Trinket)
//	if err != nil {
//		return err
//	}
//	log.Debug("Rolled new trinket", "type", trinket.Type, "name", trinket.Name)
//
//	extraRoll := d20.Roll()
//	extraType := pickFrom([]item.ItemType{item.Tool, item.Trinket})
//	if extraRoll > 17 {
//		extraType = item.SpellBook
//	} else if extraRoll > 13 {
//		extraType = pickFrom([]item.ItemType{item.Weapon, item.Armor})
//	} else if extraRoll > 5 {
//		extraType = item.Gear
//	}
//
//	extra, err := picker.PickItem(extraType)
//	if err != nil {
//		return err
//	}
//	log.Debug("Rolled new extra", "type", extra.Type, "name", extra.Name)
//
//	// Now we actually pack the inventory... fun fact, some combinations of items won't actually fit!
//	for _, thing := range []item.Item{extra, gear, tool, trinket} {
//		err := inv.AddToBackpack(thing)
//		if err != nil {
//			return err
//		}
//	}
//
//	for _, thing := range []item.Item{shield, helmet, armor, weapon} {
//		err := inv.AddActive(thing)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
