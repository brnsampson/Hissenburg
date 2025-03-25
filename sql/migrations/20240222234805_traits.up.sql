CREATE TABLE physique (
    physique_id INTEGER PRIMARY KEY,
    physique TEXT NOT NULL
);

CREATE TABLE skin (
    skin_id INTEGER PRIMARY KEY,
    skin TEXT NOT NULL
);

CREATE TABLE hair (
    hair_id INTEGER PRIMARY KEY,
    hair TEXT NOT NULL
);

CREATE TABLE face (
    face_id INTEGER PRIMARY KEY,
    face TEXT NOT NULL
);

CREATE TABLE speech (
    speech_id INTEGER PRIMARY KEY,
    speech TEXT NOT NULL
);

CREATE TABLE clothing (
    clothing_id INTEGER PRIMARY KEY,
    clothing TEXT NOT NULL
);

CREATE TABLE virtue (
    virtue_id INTEGER PRIMARY KEY,
    virtue TEXT NOT NULL
);

CREATE TABLE vice (
    vice_id INTEGER PRIMARY KEY,
    vice TEXT NOT NULL
);

CREATE TABLE reputation (
    reputation_id INTEGER PRIMARY KEY,
    reputation TEXT NOT NULL
);

CREATE TABLE misfortune (
    misfortune_id INTEGER PRIMARY KEY,
    misfortune TEXT NOT NULL
);
