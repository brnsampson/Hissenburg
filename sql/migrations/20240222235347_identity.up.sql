CREATE TABLE genders (
    id INTEGER PRIMARY KEY,
    gender TEXT NOT NULL
);

CREATE TABLE names (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE name_genders (
    id INTEGER PRIMARY KEY,
    name INTEGER NOT NULL REFERENCES name(id),
    gender INTEGER NOT NULL REFERENCES gender(id)
);

CREATE TABLE surnames (
    id INTEGER PRIMARY KEY,
    surname TEXT NOT NULL
);

CREATE TABLE backgrounds (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL
);
