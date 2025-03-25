-- name: GetEquipmentView :many
SELECT
    equipment_loadouts.id,
    equipment_loadouts.inventory,
    equipment_loadouts.slot,
    sqlc.embed(item_views)
FROM equipment_loadouts
INNER JOIN item_views ON equipment_loadouts.item = item_views.id
WHERE equipment_loadouts.id = ?
ORDER BY equipment_loadouts.slot;

-- name: GetBackpackView :many
SELECT
    items.*
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "backpack"
ORDER BY ii.container_index;

-- name: GetBonusItemView :many
SELECT
    items.*
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "bonus"
ORDER BY ii.container_index;

-- name: GetGroundView :many
SELECT
    items.*
FROM inventory_items AS ii
INNER JOIN items ON ii.item = items.id
INNER JOIN item_slots ON ii.slot = item_slots.id
WHERE ii.inventory = ?
AND item_slots.name = "ground"
ORDER BY ii.container_index;
--SELECT
--    items.*
--FROM items_on_ground
--INNER JOIN items ON items_on_ground.item = items.id
--WHERE items_on_ground.id = ?
--ORDER BY items_on_ground.container_index;

-- name: ListInventory :many
SELECT
    inventory_items.id,
    sqlc.embed(item_slots),
    inventory_items.container_index,
    sqlc.embed(item_views)
FROM inventory_items
LEFT JOIN item_slots ON inventory_items.slot = item_slots.id
LEFT JOIN item_views ON inventory_items.item = item_views.id
WHERE inventory_items.id = ?
ORDER BY inventory_items.slot, inventory_items.container_index;

-- name: AddInventoryItem :one
INSERT INTO inventory_items (inventory, slot, container_index, item) VALUES (?, ?, ?, ?) RETURNING *;

-- name: RemoveInventoryItem :exec
DELETE FROM inventory_items
WHERE inventory = ? AND slot = ? AND container_index = ?;

-- name: RemoveAllInventoryItems :exec
DELETE FROM inventory_items
WHERE inventory = ?;

-- name: UpdateInventoryItem :exec
UPDATE inventory_items
set item = ?
WHERE inventory = ? AND slot = ? AND container_index = ?;

---- name: GetInventory :one
SELECT
    inventory_items.*
FROM inventory_items AS ii
WHERE ii.inventory = ?
ORDER BY ii.item_slot, ii.container_index;
--SELECT * FROM inventories
--WHERE id = ? LIMIT 1;

-- name: GetInventoryForCharacter :one
SELECT inventories.id FROM inventories
INNER JOIN characters ON inventories.id = characters.inventory
WHERE characters.id = ? LIMIT 1;

-- name: CreateInventory :one
INSERT INTO inventories (id) VALUES (NULL) RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM inventories
WHERE id = ?;
