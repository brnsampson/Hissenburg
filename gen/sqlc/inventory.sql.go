// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: inventory.sql

package sqlc

import (
	"context"
)

const addInventoryItem = `-- name: AddInventoryItem :one
INSERT INTO inventory_items (inventory, slot, container_index, item) VALUES (?, ?, ?, ?) RETURNING id, inventory, slot, container_index, item
`

type AddInventoryItemParams struct {
	Inventory      int64
	Slot           int64
	ContainerIndex int64
	Item           int64
}

func (q *Queries) AddInventoryItem(ctx context.Context, arg AddInventoryItemParams) (InventoryItem, error) {
	row := q.db.QueryRowContext(ctx, addInventoryItem,
		arg.Inventory,
		arg.Slot,
		arg.ContainerIndex,
		arg.Item,
	)
	var i InventoryItem
	err := row.Scan(
		&i.ID,
		&i.Inventory,
		&i.Slot,
		&i.ContainerIndex,
		&i.Item,
	)
	return i, err
}

const createInventory = `-- name: CreateInventory :one
INSERT INTO inventories (id) VALUES (NULL) RETURNING id
`

func (q *Queries) CreateInventory(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, createInventory)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteInventory = `-- name: DeleteInventory :exec
DELETE FROM inventories
WHERE id = ?
`

func (q *Queries) DeleteInventory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteInventory, id)
	return err
}

const getBackpackView = `-- name: GetBackpackView :many
SELECT
    items.id, items.name, items.kind, items.slot, items.description, items.value, items.dice_count, items.dice_sides, items.armor, items.storage, items.size, items.active_size, items.stackable, items.icon
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "backpack"
ORDER BY ii.container_index
`

func (q *Queries) GetBackpackView(ctx context.Context, inventory int64) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getBackpackView, inventory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Kind,
			&i.Slot,
			&i.Description,
			&i.Value,
			&i.DiceCount,
			&i.DiceSides,
			&i.Armor,
			&i.Storage,
			&i.Size,
			&i.ActiveSize,
			&i.Stackable,
			&i.Icon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBonusItemView = `-- name: GetBonusItemView :many
SELECT
    items.id, items.name, items.kind, items.slot, items.description, items.value, items.dice_count, items.dice_sides, items.armor, items.storage, items.size, items.active_size, items.stackable, items.icon
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "bonus"
ORDER BY ii.container_index
`

func (q *Queries) GetBonusItemView(ctx context.Context, inventory int64) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getBonusItemView, inventory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Kind,
			&i.Slot,
			&i.Description,
			&i.Value,
			&i.DiceCount,
			&i.DiceSides,
			&i.Armor,
			&i.Storage,
			&i.Size,
			&i.ActiveSize,
			&i.Stackable,
			&i.Icon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEquipmentView = `-- name: GetEquipmentView :many
SELECT
    equipment_loadouts.id,
    equipment_loadouts.inventory,
    equipment_loadouts.slot,
    item_views.id, item_views.name, item_views.kind, item_views.slot, item_views.description, item_views.value, item_views.dice_count, item_views.dice_sides, item_views.armor, item_views.storage, item_views.size, item_views.active_size, item_views.stackable, item_views.icon
FROM equipment_loadouts
INNER JOIN item_views ON equipment_loadouts.item = item_views.id
WHERE equipment_loadouts.id = ?
ORDER BY equipment_loadouts.slot
`

type GetEquipmentViewRow struct {
	ID        int64
	Inventory int64
	Slot      string
	ItemView  ItemView
}

func (q *Queries) GetEquipmentView(ctx context.Context, id int64) ([]GetEquipmentViewRow, error) {
	rows, err := q.db.QueryContext(ctx, getEquipmentView, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEquipmentViewRow
	for rows.Next() {
		var i GetEquipmentViewRow
		if err := rows.Scan(
			&i.ID,
			&i.Inventory,
			&i.Slot,
			&i.ItemView.ID,
			&i.ItemView.Name,
			&i.ItemView.Kind,
			&i.ItemView.Slot,
			&i.ItemView.Description,
			&i.ItemView.Value,
			&i.ItemView.DiceCount,
			&i.ItemView.DiceSides,
			&i.ItemView.Armor,
			&i.ItemView.Storage,
			&i.ItemView.Size,
			&i.ItemView.ActiveSize,
			&i.ItemView.Stackable,
			&i.ItemView.Icon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroundView = `-- name: GetGroundView :many
SELECT
    items.id, items.name, items.kind, items.slot, items.description, items.value, items.dice_count, items.dice_sides, items.armor, items.storage, items.size, items.active_size, items.stackable, items.icon
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "ground"
ORDER BY ii.container_index
`

func (q *Queries) GetGroundView(ctx context.Context, inventory int64) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getGroundView, inventory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Kind,
			&i.Slot,
			&i.Description,
			&i.Value,
			&i.DiceCount,
			&i.DiceSides,
			&i.Armor,
			&i.Storage,
			&i.Size,
			&i.ActiveSize,
			&i.Stackable,
			&i.Icon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInventoryForCharacter = `-- name: GetInventoryForCharacter :one

SELECT inventories.id FROM inventories
INNER JOIN characters ON inventories.id = characters.inventory
WHERE characters.id = ? LIMIT 1
`

// SELECT * FROM inventories
// WHERE id = ? LIMIT 1;
func (q *Queries) GetInventoryForCharacter(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getInventoryForCharacter, id)
	err := row.Scan(&id)
	return id, err
}

const listInventory = `-- name: ListInventory :many

SELECT
    inventory_items.id,
    item_slots.id, item_slots.name,
    inventory_items.container_index,
    item_views.id, item_views.name, item_views.kind, item_views.slot, item_views.description, item_views.value, item_views.dice_count, item_views.dice_sides, item_views.armor, item_views.storage, item_views.size, item_views.active_size, item_views.stackable, item_views.icon
FROM inventory_items
LEFT JOIN item_slots ON inventory_items.slot = item_slots.id
LEFT JOIN item_views ON inventory_items.item = item_views.id
WHERE inventory_items.id = ?
ORDER BY inventory_items.slot, inventory_items.container_index
`

type ListInventoryRow struct {
	ID             int64
	ItemSlot       ItemSlot
	ContainerIndex int64
	ItemView       ItemView
}

// SELECT
//
//	items.*
//
// FROM items_on_ground
// INNER JOIN items ON items_on_ground.item = items.id
// WHERE items_on_ground.id = ?
// ORDER BY items_on_ground.container_index;
func (q *Queries) ListInventory(ctx context.Context, id int64) ([]ListInventoryRow, error) {
	rows, err := q.db.QueryContext(ctx, listInventory, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListInventoryRow
	for rows.Next() {
		var i ListInventoryRow
		if err := rows.Scan(
			&i.ID,
			&i.ItemSlot.ID,
			&i.ItemSlot.Name,
			&i.ContainerIndex,
			&i.ItemView.ID,
			&i.ItemView.Name,
			&i.ItemView.Kind,
			&i.ItemView.Slot,
			&i.ItemView.Description,
			&i.ItemView.Value,
			&i.ItemView.DiceCount,
			&i.ItemView.DiceSides,
			&i.ItemView.Armor,
			&i.ItemView.Storage,
			&i.ItemView.Size,
			&i.ItemView.ActiveSize,
			&i.ItemView.Stackable,
			&i.ItemView.Icon,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeAllInventoryItems = `-- name: RemoveAllInventoryItems :exec
DELETE FROM inventory_items
WHERE inventory = ?
`

func (q *Queries) RemoveAllInventoryItems(ctx context.Context, inventory int64) error {
	_, err := q.db.ExecContext(ctx, removeAllInventoryItems, inventory)
	return err
}

const removeInventoryItem = `-- name: RemoveInventoryItem :exec
DELETE FROM inventory_items
WHERE inventory = ? AND slot = ? AND container_index = ?
`

type RemoveInventoryItemParams struct {
	Inventory      int64
	Slot           int64
	ContainerIndex int64
}

func (q *Queries) RemoveInventoryItem(ctx context.Context, arg RemoveInventoryItemParams) error {
	_, err := q.db.ExecContext(ctx, removeInventoryItem, arg.Inventory, arg.Slot, arg.ContainerIndex)
	return err
}

const updateInventoryItem = `-- name: UpdateInventoryItem :exec
UPDATE inventory_items
set item = ?
WHERE inventory = ? AND slot = ? AND container_index = ?
`

type UpdateInventoryItemParams struct {
	Item           int64
	Inventory      int64
	Slot           int64
	ContainerIndex int64
}

func (q *Queries) UpdateInventoryItem(ctx context.Context, arg UpdateInventoryItemParams) error {
	_, err := q.db.ExecContext(ctx, updateInventoryItem,
		arg.Item,
		arg.Inventory,
		arg.Slot,
		arg.ContainerIndex,
	)
	return err
}
