CREATE TABLE item_kinds (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE item_slots (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE items (
    id   INTEGER PRIMARY KEY,
    name TEXT COLLATE NOCASE UNIQUE NOT NULL,
    kind INTEGER NOT NULL REFERENCES item_kinds(id),
    slot INTEGER NOT NULL REFERENCES item_slots(id),
    description TEXT NOT NULL,
    value INTEGER NOT NULL,
    dice_count TINYINT NOT NULL,
    dice_sides TINYINT NOT NULL,
    armor TINYINT NOT NULL,
    storage TINYINT NOT NULL,
    size TINYINT NOT NULL,
    active_size TINYINT NOT NULL,
    stackable BOOL NOT NULL,
    icon TEXT NOT NULL,
    UNIQUE (name, kind)
);

CREATE VIEW item_views AS
SELECT
    items.id,
    items.name,
    item_kinds.name AS kind,
    item_slots.name AS slot,
    items.description,
    items.value,
    items.dice_count,
    items.dice_sides,
    items.armor,
    items.storage,
    items.size,
    items.active_size,
    items.stackable,
    items.icon
FROM items
INNER JOIN item_kinds ON items.kind = item_kinds.id
INNER JOIN item_slots ON items.slot = item_slots.id
