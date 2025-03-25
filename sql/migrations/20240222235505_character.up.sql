CREATE TABLE parties (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    inventory INTEGER REFERENCES inventories(id)
);

CREATE TABLE characters (
    id INTEGER PRIMARY KEY,
    description TEXT NOT NULL,
    -- Ownership and Associations
    user INTEGER NOT NULL REFERENCES user(id),
    party INTEGER NOT NULL REFERENCES parties(id),
    village INTEGER NOT NULL REFERENCES village(id),
    -- Identity
    gender TEXT NOT NULL,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    age INTEGER NOT NULL,
    portrait TEXT NOT NULL,
    background INTEGER NOT NULL REFERENCES background(id),
    -- Traits
    physique TEXT NOT NULL,
    skin TEXT NOT NULL,
    hair TEXT NOT NULL,
    face TEXT NOT NULL,
    speech TEXT NOT NULL,
    clothing TEXT NOT NULL,
    virtue TEXT NOT NULL,
    vice TEXT NOT NULL,
    reputation TEXT NOT NULL,
    misfortune TEXT NOT NULL,
    -- Status
    hp TINYINT NOT NULL,
    max_hp TINYINT NOT NULL,
    str TINYINT NOT NULL,
    max_str TINYINT NOT NULL,
    dex TINYINT NOT NULL,
    max_dex TINYINT NOT NULL,
    will TINYINT NOT NULL,
    max_will TINYINT NOT NULL,
    -- Inventory
    inventory INTEGER NOT NULL REFERENCES inventories(id),
    UNIQUE (user, name, surname)
);
