package data

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/brnsampson/Hissenburg/models/inventory"
	"github.com/brnsampson/Hissenburg/gen/sqlc"
)

type InventoryRepo interface {
	ItemRepo
	GetEquipmentView(ctx context.Context, id int64) (inventory.Equipment, error)
	GetBackpackView(ctx context.Context, id int64) ([]sqlc.Item, error)
	GetBonusItemView(ctx context.Context, id int64) ([]sqlc.Item, error)
	GetGroundView(ctx context.Context, id int64) ([]sqlc.Item, error)
	AddInventoryItem(ctx context.Context, inventory int64, slot, index, item int64) (sqlc.InventoryItem, error)
	RemoveInventoryItem(ctx context.Context, inventory int64, slot, index int64) error
	RemoveAllInventoryItems(ctx context.Context, inventory int64) error
	UpdateInventoryItem(ctx context.Context, inventory int64, slot, index, item int64) error
	GetInventory(ctx context.Context, id int64) (inventory.Inventory, error)
	GetInventoryForCharacter(ctx context.Context, id int64) (int64, error)
	CreateInventory(ctx context.Context) (int64, error)
	DeleteInventory(ctx context.Context, id int64) error
	UpdateInventory(ctx context.Context, inv inventory.Inventory) error
}

func (r *Repository) AddInventoryItem(ctx context.Context, inventory int64, slot, index, item int64) (sqlc.InventoryItem, error) {
	params := sqlc.AddInventoryItemParams{Inventory: inventory, Slot: slot, ContainerIndex: index, Item: item}
	return r.Queries.AddInventoryItem(ctx, params)
}

func (r *Repository) RemoveInventoryItem(ctx context.Context, inventory int64, slot, index int64) error {
	params := sqlc.RemoveInventoryItemParams{Inventory: inventory, Slot: slot, ContainerIndex: index}
	return r.Queries.RemoveInventoryItem(ctx, params)
}

func (r *Repository) UpdateInventoryItem(ctx context.Context, inventory int64, slot, index, item int64) error {
	params := sqlc.UpdateInventoryItemParams{Inventory: inventory, Slot: slot, ContainerIndex: index, Item: item}
	return r.Queries.UpdateInventoryItem(ctx, params)
}

func (r *Repository) GetEquipmentView(ctx context.Context, id int64) (inventory.Equipment, error) {
	equip := inventory.Equipment{ }
	items, err := r.Queries.GetEquipmentView(ctx, id)
	if err != nil {
		log.Warn("Failed to look up equipment for inventory", "inventory", id)
		return equip, err
	}
	for _, row := range items {
		if row.Slot == "head" {
			equip.LeftHand = &row.ItemView
		} else if row.Slot == "right_hand" {
			equip.RightHand = &row.ItemView
		} else if row.Slot == "head" {
			equip.Head = &row.ItemView
		} else if row.Slot == "torso" {
			equip.Torso = &row.ItemView
		} else {
			log.Warn("Encountered unknown item slot in inventory equipment... Ignoring.", "slot", row.Slot)
		}
	}

	return equip, nil
}

func (r *Repository) GetInventory(ctx context.Context, id int64) (inventory.Inventory, error) {
	inv := inventory.New()
	inv.ID = id
	items, err := r.Queries.ListInventory(ctx, id)
	if err != nil {
		log.Warn("Failed to look up items for inventory", "inventory", id)
		return inv, err
	}
	if len(items) < 1 {
		log.Debug("No items in inventory", "inventory", id)
		return inv, fmt.Errorf("No Items in Inventory")
	}

	for _, row := range items {
		if row.ItemSlot.Name == "backpack" {
			inv.InsertIntoBackpack(inventory.ContainerItem{ Index: int(row.ContainerIndex), Item: row.ItemView })
		} else if row.ItemSlot.Name == "bonus" {
			inv.InsertIntoBonusSpace(inventory.ContainerItem{ Index: int(row.ContainerIndex), Item: row.ItemView })
		} else if row.ItemSlot.Name == "ground" {
			inv.InsertIntoGround(inventory.ContainerItem{ Index: int(row.ContainerIndex), Item: row.ItemView })
		} else if row.ItemSlot.Name == "left_hand" {
			inv.Equipment.LeftHand = &row.ItemView
		} else if row.ItemSlot.Name == "right_hand" {
			inv.Equipment.RightHand = &row.ItemView
		} else if row.ItemSlot.Name == "head" {
			inv.Equipment.Head = &row.ItemView
		} else if row.ItemSlot.Name == "torso" {
			inv.Equipment.Torso = &row.ItemView
		} else {
			log.Warn("Encountered unknown item slot in inventory... Ignoring.", "slot", row.ItemSlot.Name)
		}
	}

	return inv, nil
}

func (r *Repository) UpdateInventory(ctx context.Context, inv inventory.Inventory) error {
	dbInv, err := r.GetInventory(ctx, inv.ID)
	if err != nil {
		log.Warn("Failed to look up items for inventory", "inventory", inv.ID)
		return err
	}

	toAdd, toRemove, updates := inventory.DiffInventory(inv, dbInv)

	slotName := "left_hand"
	leftSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "right_hand"
	rightSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "head"
	headSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "torso"
	torsoSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "backpack"
	backpackSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "bonus"
	bonusSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	slotName = "ground"
	groundSlot, err := r.GetItemSlotFromString(ctx, slotName)
	if err != nil {
		log.Error("Failed to update inventory: failed to look up expected item slot", "inventory", inv.ID, "slot", slotName)
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Error("Failed to initialize transaction to update inventory!")
		return err
	}
	defer tx.Rollback()
	qtx := r.Queries.WithTx(tx)

	// Update all in toUpdate
	if updates.Equipment.LeftHand != nil {
		slot := leftSlot.ID
		toUpdate := updates.Equipment.LeftHand.Left
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0, Item: toUpdate.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if updates.Equipment.RightHand != nil {
		slot := rightSlot.ID
		toUpdate := updates.Equipment.RightHand.Left
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0, Item: toUpdate.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if updates.Equipment.Head != nil {
		slot := headSlot.ID
		toUpdate := updates.Equipment.Head.Left
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0, Item: toUpdate.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if updates.Equipment.Torso != nil {
		slot := torsoSlot.ID
		toUpdate := updates.Equipment.Torso.Left
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0, Item: toUpdate.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	for _, update := range updates.Backpack {
		slot := backpackSlot.ID
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(update.Index), Item: update.Left.Item.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot, "index", update.Index)
			return err
		}
	}
	for _, update := range updates.BonusSpace {
		slot := bonusSlot.ID
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(update.Index), Item: update.Left.Item.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot, "index", update.Index)
			return err
		}
	}
	for _, update := range updates.Ground {
		slot := groundSlot.ID
		params := sqlc.UpdateInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(update.Index), Item: update.Left.Item.ID }
		err := qtx.UpdateInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to update inventory", "inventory", inv.ID, "slot", slot, "index", update.Index)
			return err
		}
	}

	// Remove all in toRemove
	if toRemove.Equipment.LeftHand != nil {
		slot := leftSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toRemove.Equipment.RightHand != nil {
		slot := rightSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toRemove.Equipment.Head != nil {
		slot := headSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toRemove.Equipment.Torso != nil {
		slot := torsoSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	for _, rm := range toRemove.Backpack {
		slot := backpackSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(rm.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", rm.Index)
			return err
		}
	}
	for _, rm := range toRemove.BonusSpace {
		slot := bonusSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(rm.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", rm.Index)
			return err
		}
	}
	for _, rm := range toRemove.Ground {
		slot := groundSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(rm.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", rm.Index)
			return err
		}
	}

	// Add all in toAdd
	if toAdd.Equipment.LeftHand != nil {
		slot := leftSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toAdd.Equipment.RightHand != nil {
		slot := rightSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toAdd.Equipment.Head != nil {
		slot := headSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	if toAdd.Equipment.Torso != nil {
		slot := torsoSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: 0 }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot)
			return err
		}
	}
	for _, add := range toAdd.Backpack {
		slot := backpackSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(add.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", add.Index)
			return err
		}
	}
	for _, add := range toAdd.BonusSpace {
		slot := bonusSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(add.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", add.Index)
			return err
		}
	}
	for _, add := range toAdd.Ground {
		slot := groundSlot.ID
		params := sqlc.RemoveInventoryItemParams { Inventory: inv.ID, Slot: slot, ContainerIndex: int64(add.Index) }
		err := qtx.RemoveInventoryItem(ctx, params)
		if err != nil {
			log.Error("Failed to remove item from inventory", "inventory", inv.ID, "slot", slot, "index", add.Index)
			return err
		}
	}

	return tx.Commit()
}

//func (r *Repository) CreateInventory(ctx context.Context, inv sqlc.Inventory) (sqlc.Inventory, error) {
//	invParams := sqlc.CreateInventoryParams {
//		LeftHand: inv.LeftHand,
//		RightHand: inv.RightHand,
//		Head: inv.Head,
//		Torso: inv.Torso,
//		Backpack0: inv.Backpack0,
//		Backpack1: inv.Backpack1,
//		Backpack2: inv.Backpack2,
//		Backpack3: inv.Backpack3,
//		Backpack4: inv.Backpack4,
//		Backpack5: inv.Backpack5,
//		ExtraSpace0: inv.ExtraSpace0,
//		ExtraSpace1: inv.ExtraSpace1,
//		ExtraSpace2: inv.ExtraSpace2,
//		ExtraSpace3: inv.ExtraSpace3,
//		ExtraSpace4: inv.ExtraSpace4,
//		ExtraSpace5: inv.ExtraSpace5,
//		Ground0: inv.Ground0,
//		Ground1: inv.Ground1,
//		Ground2: inv.Ground2,
//		Ground3: inv.Ground3,
//		Ground4: inv.Ground4,
//		Ground5: inv.Ground5,
//	}
//	return r.Queries.CreateInventory(ctx, invParams)
//}

//func (r *Repository) CreateInventory(ctx context.Context, inv inventory.Inventory) (inventory.Inventory, error) {
//	invalid, err := r.Queries.GetItemFromName(ctx, "MissingNo.")
//	if err != nil {
//		log.Error("Error Creating Inventory: Failed to lookup Invalid Item", "error", err)
//		return inv, err
//	}
//
//	left, err := r.Queries.GetItemFromName(ctx, inv.LeftHand.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		left = invalid
//	}
//
//	right, err := r.Queries.GetItemFromName(ctx, inv.RightHand.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		right = invalid
//	}
//
//	head, err := r.Queries.GetItemFromName(ctx, inv.Head.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		head = invalid
//	}
//
//	torso, err := r.Queries.GetItemFromName(ctx, inv.Torso.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		torso = invalid
//	}
//
//	backpack0, err := r.Queries.GetItemFromName(ctx, inv.Backpack0.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack0 = invalid
//	}
//
//	backpack1, err := r.Queries.GetItemFromName(ctx, inv.Backpack1.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack1 = invalid
//	}
//
//	backpack2, err := r.Queries.GetItemFromName(ctx, inv.Backpack2.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack2 = invalid
//	}
//
//	backpack3, err := r.Queries.GetItemFromName(ctx, inv.Backpack3.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack3 = invalid
//	}
//
//	backpack4, err := r.Queries.GetItemFromName(ctx, inv.Backpack4.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack4 = invalid
//	}
//
//	backpack5, err := r.Queries.GetItemFromName(ctx, inv.Backpack5.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		backpack5 = invalid
//	}
//
//	extra0, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace0.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra0 = invalid
//	}
//
//	extra1, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace1.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra1 = invalid
//	}
//
//	extra2, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace2.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra2 = invalid
//	}
//
//	extra3, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace3.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra3 = invalid
//	}
//
//	extra4, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace4.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra4 = invalid
//	}
//
//	extra5, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace5.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		extra5 = invalid
//	}
//
//	ground0, err := r.Queries.GetItemFromName(ctx, inv.Ground0.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground0 = invalid
//	}
//
//	ground1, err := r.Queries.GetItemFromName(ctx, inv.Ground1.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground1 = invalid
//	}
//
//	ground2, err := r.Queries.GetItemFromName(ctx, inv.Ground2.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground2 = invalid
//	}
//
//	ground3, err := r.Queries.GetItemFromName(ctx, inv.Ground3.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground3 = invalid
//	}
//
//	ground4, err := r.Queries.GetItemFromName(ctx, inv.Ground4.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground4 = invalid
//	}
//
//	ground5, err := r.Queries.GetItemFromName(ctx, inv.Ground5.Name)
//	if err != nil {
//		log.Warn("Invalid item encountered when creating Inventory", "error", err)
//		ground5 = invalid
//	}
//
//	invParams := sqlc.CreateInventoryParams {
//		LeftHand: left.ItemID,
//		RightHand: right.ItemID,
//		Head: head.ItemID,
//		Torso: torso.ItemID,
//		Backpack0: backpack0.ItemID,
//		Backpack1: backpack1.ItemID,
//		Backpack2: backpack2.ItemID,
//		Backpack3: backpack3.ItemID,
//		Backpack4: backpack4.ItemID,
//		Backpack5: backpack5.ItemID,
//		ExtraSpace0: extra0.ItemID,
//		ExtraSpace1: extra1.ItemID,
//		ExtraSpace2: extra2.ItemID,
//		ExtraSpace3: extra3.ItemID,
//		ExtraSpace4: extra4.ItemID,
//		ExtraSpace5: extra5.ItemID,
//		Ground0: ground0.ItemID,
//		Ground1: ground1.ItemID,
//		Ground2: ground2.ItemID,
//		Ground3: ground3.ItemID,
//		Ground4: ground4.ItemID,
//		Ground5: ground5.ItemID,
//	}
//
//	newInv, err := r.Queries.CreateInventory(ctx, invParams)
//	if err != nil {
//		log.Error("Failed to create Inventory", "inventory", inv, "error", err)
//		return inv, err
//	}
//
//	return InventoryDBToModel(ctx, r.Queries, newInv)
//}

//func (r *Repository) UpdateInventory(ctx context.Context, inv inventory.Inventory) error {
//	invalid, err := r.Queries.GetItemFromName(ctx, "MissingNo.")
//	if err != nil {
//		return err
//	}
//
//	// Just to check that the inventory already exists
//	_, err = r.Queries.GetInventory(ctx, inv.ID)
//	if err != nil {
//		return err
//	}
//
//	left, err := r.Queries.GetItemFromName(ctx, inv.LeftHand.Name)
//	if err != nil {
//		left = invalid
//	}
//
//	right, err := r.Queries.GetItemFromName(ctx, inv.RightHand.Name)
//	if err != nil {
//		right = invalid
//	}
//
//	head, err := r.Queries.GetItemFromName(ctx, inv.Head.Name)
//	if err != nil {
//		head = invalid
//	}
//
//	torso, err := r.Queries.GetItemFromName(ctx, inv.Torso.Name)
//	if err != nil {
//		torso = invalid
//	}
//
//	backpack0, err := r.Queries.GetItemFromName(ctx, inv.Backpack0.Name)
//	if err != nil {
//		backpack0 = invalid
//	}
//
//	backpack1, err := r.Queries.GetItemFromName(ctx, inv.Backpack1.Name)
//	if err != nil {
//		backpack1 = invalid
//	}
//
//	backpack2, err := r.Queries.GetItemFromName(ctx, inv.Backpack2.Name)
//	if err != nil {
//		backpack2 = invalid
//	}
//
//	backpack3, err := r.Queries.GetItemFromName(ctx, inv.Backpack3.Name)
//	if err != nil {
//		backpack3 = invalid
//	}
//
//	backpack4, err := r.Queries.GetItemFromName(ctx, inv.Backpack4.Name)
//	if err != nil {
//		backpack4 = invalid
//	}
//
//	backpack5, err := r.Queries.GetItemFromName(ctx, inv.Backpack5.Name)
//	if err != nil {
//		backpack5 = invalid
//	}
//
//	extra0, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace0.Name)
//	if err != nil {
//		extra0 = invalid
//	}
//
//	extra1, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace1.Name)
//	if err != nil {
//		extra1 = invalid
//	}
//
//	extra2, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace2.Name)
//	if err != nil {
//		extra2 = invalid
//	}
//
//	extra3, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace3.Name)
//	if err != nil {
//		extra3 = invalid
//	}
//
//	extra4, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace4.Name)
//	if err != nil {
//		extra4 = invalid
//	}
//
//	extra5, err := r.Queries.GetItemFromName(ctx, inv.ExtraSpace5.Name)
//	if err != nil {
//		extra5 = invalid
//	}
//
//	ground0, err := r.Queries.GetItemFromName(ctx, inv.Ground0.Name)
//	if err != nil {
//		ground0 = invalid
//	}
//
//	ground1, err := r.Queries.GetItemFromName(ctx, inv.Ground1.Name)
//	if err != nil {
//		ground1 = invalid
//	}
//
//	ground2, err := r.Queries.GetItemFromName(ctx, inv.Ground2.Name)
//	if err != nil {
//		ground2 = invalid
//	}
//
//	ground3, err := r.Queries.GetItemFromName(ctx, inv.Ground3.Name)
//	if err != nil {
//		ground3 = invalid
//	}
//
//	ground4, err := r.Queries.GetItemFromName(ctx, inv.Ground4.Name)
//	if err != nil {
//		ground4 = invalid
//	}
//
//	ground5, err := r.Queries.GetItemFromName(ctx, inv.Ground5.Name)
//	if err != nil {
//		ground5 = invalid
//	}
//
//	invParams := sqlc.UpdateInventoryParams {
//		LeftHand: left.ItemID,
//		RightHand: right.ItemID,
//		Head: head.ItemID,
//		Torso: torso.ItemID,
//		Backpack0: backpack0.ItemID,
//		Backpack1: backpack1.ItemID,
//		Backpack2: backpack2.ItemID,
//		Backpack3: backpack3.ItemID,
//		Backpack4: backpack4.ItemID,
//		Backpack5: backpack5.ItemID,
//		ExtraSpace0: extra0.ItemID,
//		ExtraSpace1: extra1.ItemID,
//		ExtraSpace2: extra2.ItemID,
//		ExtraSpace3: extra3.ItemID,
//		ExtraSpace4: extra4.ItemID,
//		ExtraSpace5: extra5.ItemID,
//		Ground0: ground0.ItemID,
//		Ground1: ground1.ItemID,
//		Ground2: ground2.ItemID,
//		Ground3: ground3.ItemID,
//		Ground4: ground4.ItemID,
//		Ground5: ground5.ItemID,
//		InventoryID: inv.ID,
//	}
//
//	return r.Queries.UpdateInventory(ctx, invParams)
//}
//
//
