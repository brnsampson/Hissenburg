--CREATE TABLE backpacks (
--    id INTEGER PRIMARY KEY,
--    slot_0 INTEGER NOT NULL REFERENCES items(id),
--    slot_1 INTEGER NOT NULL REFERENCES items(id),
--    slot_2 INTEGER NOT NULL REFERENCES items(id),
--    slot_3 INTEGER NOT NULL REFERENCES items(id),
--    slot_4 INTEGER NOT NULL REFERENCES items(id),
--    slot_5 INTEGER NOT NULL REFERENCES items(id)
--);
--
--CREATE TABLE extra_spaces (
--    id INTEGER PRIMARY KEY,
--    slot_0 INTEGER NOT NULL REFERENCES items(id),
--    slot_1 INTEGER NOT NULL REFERENCES items(id),
--    slot_2 INTEGER NOT NULL REFERENCES items(id),
--    slot_3 INTEGER NOT NULL REFERENCES items(id),
--    slot_4 INTEGER NOT NULL REFERENCES items(id),
--    slot_5 INTEGER NOT NULL REFERENCES items(id)
--);
--
--CREATE TABLE ground_drops (
--    id INTEGER PRIMARY KEY,
--    slot_0 INTEGER NOT NULL REFERENCES items(id),
--    slot_1 INTEGER NOT NULL REFERENCES items(id),
--    slot_2 INTEGER NOT NULL REFERENCES items(id),
--    slot_3 INTEGER NOT NULL REFERENCES items(id),
--    slot_4 INTEGER NOT NULL REFERENCES items(id),
--    slot_5 INTEGER NOT NULL REFERENCES items(id)
--);
--
--CREATE TABLE equipment_loadouts (
--    id INTEGER PRIMARY KEY,
--    left_hand INTEGER NOT NULL REFERENCES items(id),
--    right_hand INTEGER NOT NULL REFERENCES items(id),
--    head INTEGER NOT NULL REFERENCES items(id),
--    torso INTEGER NOT NULL REFERENCES items(id)
--);
--
--CREATE TABLE inventories (
--    id INTEGER PRIMARY KEY,
--    equipment INTEGER NOT NULL REFERENCES equipment_loadouts(id),
--    backpack INTEGER NOT NULL REFERENCES backpacks(id),
--    extra_space INTEGER NOT NULL REFERENCES extra_spaces(id),
--    ground INTEGER NOT NULL REFERENCES ground_drop(id)
--);

CREATE TABLE inventory_items (
    id INTEGER NOT NULL PRIMARY KEY,
    inventory INTEGER NOT NULL REFERENCES inventories(id),
    slot INTEGER NOT NULL REFERENCES item_slot(id),
    container_index INTEGER NOT NULL,
    item INTEGER NOT NULL REFERENCES items(id)
);

CREATE TABLE inventories (
    id INTEGER PRIMARY KEY
);

CREATE VIEW backpacks AS
SELECT
    inventory_items.id,
    inventories.id AS inventory,
    item_slots.name AS slot,
    inventory_items.container_index,
    inventory_items.item
FROM inventories
INNER JOIN inventory_items ON inventories.id = inventory_items.inventory
INNER JOIN item_slots ON inventory_items.slot = item_slots.id
WHERE item_slots.name = "backpack";

CREATE VIEW bonus_slots AS
SELECT
    inventory_items.id,
    inventories.id AS inventory,
    item_slots.name AS slot,
    inventory_items.container_index,
    inventory_items.item
FROM inventories
INNER JOIN inventory_items ON inventories.id = inventory_items.inventory
INNER JOIN item_slots ON inventory_items.slot = item_slots.id
WHERE item_slots.name = "bonus";

CREATE VIEW items_on_ground AS
SELECT
    inventory_items.id,
    inventories.id AS inventory,
    item_slots.name AS slot,
    inventory_items.container_index,
    inventory_items.item
FROM inventories
INNER JOIN inventory_items ON inventories.id = inventory_items.inventory
INNER JOIN item_slots ON inventory_items.slot = item_slots.id
WHERE item_slots.name = "ground";

CREATE VIEW equipment_loadouts AS
SELECT
    inventory_items.id,
    inventories.id AS inventory,
    islot.name AS slot,
    inventory_items.container_index,
    inventory_items.item
FROM inventories
INNER JOIN inventory_items ON inventories.id = inventory_items.inventory
INNER JOIN item_slots AS islot ON inventory_items.slot = islot.id
WHERE islot.name = "head" OR islot.name = "torso" OR islot.name = "body" OR islot.name = "hands" OR islot.name = "left_hand" OR islot.name = "right_hand";

--CREATE VIEW extra_spaces AS
--SELECT
--    id,
--    extra_space_0 AS slot_0,
--    extra_space_1 AS slot_1,
--    extra_space_2 AS slot_2,
--    extra_space_3 AS slot_3,
--    extra_space_4 AS slot_4,
--    extra_space_5 AS slot_5
--FROM inventories;
--
--CREATE VIEW ground_drops AS
--SELECT
--    id,
--    ground_0 AS slot_0,
--    ground_1 AS slot_1,
--    ground_2 AS slot_2,
--    ground_3 AS slot_3,
--    ground_4 AS slot_4,
--    ground_5 AS slot_5
--FROM inventories;
--
--CREATE VIEW equipment_loadouts AS
--SELECT
--    id,
--    left_hand,
--    right_hand,
--    head,
--    torso
--FROM inventories;
