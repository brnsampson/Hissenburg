-- name: GetItemSlot :one
SELECT * FROM item_slots
WHERE id = ? LIMIT 1;

-- name: GetItemSlotFromString :one
SELECT * FROM item_slots
WHERE name = ? LIMIT 1;

-- name: ListItemSlots :many
SELECT * FROM item_slots
ORDER BY name;

-- name: CreateItemSlot :one
INSERT INTO item_slots (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteItemSlot :exec
DELETE FROM item_slots
WHERE id = ?;

-- name: GetItemKind :one
SELECT * FROM item_kinds
WHERE id = ? LIMIT 1;

-- name: GetItemKindFromString :one
SELECT * FROM item_kinds
WHERE name = ? LIMIT 1;

-- name: ListItemKinds :many
SELECT * FROM item_kinds
ORDER BY name;

-- name: CreateItemKind :one
INSERT INTO item_kinds (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteItemKind :exec
DELETE FROM item_kinds
WHERE id = ?;

-- name: GetItem :one
SELECT * FROM items
WHERE id = ? LIMIT 1;

-- name: GetItemView :one
SELECT * FROM item_views
WHERE id = ? LIMIT 1;

-- name: GetItemFromName :one
SELECT * FROM item_views
WHERE name = ? LIMIT 1;

-- name: GetItemFromKindAndName :one
SELECT * FROM item_views
WHERE kind = LOWER(sqlc.arg(kind)) AND name = Lower(sqlc.arg(name)) LIMIT 1;
--SELECT item.* FROM item
--JOIN item_kinds ON item.kind = item_kinds.item_kinds_id
--JOIN item_slots ON item.slot = item_slots.item_slots_id
--WHERE item_kinds.item_kinds = ? AND name = ? LIMIT 1;

-- name: GetRandomItemFromKind :one
SELECT * FROM item_views
WHERE kind = LOWER(sqlc.arg(kind))
ORDER BY RANDOM() LIMIT 1;
--SELECT item.* FROM item
--JOIN item_kinds ON item.kind = item_kinds.item_kinds_id
--WHERE item_kinds.item_kinds = ?
--ORDER BY RANDOM() LIMIT 1;

-- name: ListItems :many
SELECT * FROM item_views
ORDER BY kind, name;

-- name: ListItemsForKind :many
SELECT * FROM item_views
WHERE kind = LOWER(sqlc.arg(kind))
ORDER BY name;

-- name: CreateItem :one
INSERT INTO items (
    name, kind, slot, description, value, dice_count, dice_sides, armor, storage, size, active_size, stackable, icon
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateItem :exec
UPDATE items
set name = ?,
kind = ?,
slot = ?,
description = ?,
value = ?,
dice_count = ?,
dice_sides = ?,
armor = ?,
storage = ?,
size = ?,
active_size = ?,
stackable = ?,
icon = ?
WHERE id = ?;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = ?;
