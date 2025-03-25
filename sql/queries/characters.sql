-- name: GetCharacterView :one
SELECT
    characters.id,
	sqlc.embed(users),
	sqlc.embed(parties),
	sqlc.embed(villages),
    characters.gender,
    characters.name,
    characters.surname,
    characters.age,
    characters.portrait,
    sqlc.embed(backgrounds),
    characters.physique,
    characters.skin,
    characters.hair,
    characters.face,
    characters.speech,
    characters.clothing,
    characters.virtue,
    characters.vice,
    characters.reputation,
    characters.misfortune,
    characters.description,
    characters.hp,
    characters.max_hp,
    characters.str,
    characters.max_str,
    characters.dex,
    characters.max_dex,
    characters.will,
    characters.max_will,
    sqlc.embed(inventories)
FROM characters
JOIN parties ON party = parties.id
JOIN users ON user = users.id
JOIN villages ON village = villages.id
JOIN backgrounds ON background = backgrounds.id
JOIN inventories ON characters.inventory = inventories.id
WHERE characters.id = ? LIMIT 1;

-- name: GetCharacterViewFromName :one
SELECT
    characters.id,
	sqlc.embed(users),
	sqlc.embed(parties),
	sqlc.embed(villages),
    characters.gender,
    characters.name,
    characters.surname,
    characters.age,
    characters.portrait,
    sqlc.embed(backgrounds),
    characters.physique,
    characters.skin,
    characters.hair,
    characters.face,
    characters.speech,
    characters.clothing,
    characters.virtue,
    characters.vice,
    characters.reputation,
    characters.misfortune,
    characters.description,
    characters.hp,
    characters.max_hp,
    characters.str,
    characters.max_str,
    characters.dex,
    characters.max_dex,
    characters.will,
    characters.max_will,
    sqlc.embed(inventories)
FROM characters
JOIN parties ON party = parties.id
JOIN users ON user = users.id
JOIN villages ON village = villages.id
JOIN backgrounds ON background = backgrounds.id
JOIN inventories ON characters.inventory = inventories.id
WHERE characters.name = ? AND characters.surname = ? LIMIT 1;

-- name: GetCharacter :one
SELECT * FROM characters
WHERE id = ? LIMIT 1;

-- name: GetCharacterFromName :one
SELECT * FROM characters
WHERE  name = ? AND surname = ? LIMIT 1;

-- name: ListCharacters :many
SELECT * FROM characters
ORDER BY village, party, user;

-- name: ListCharactersFromVillage :many
SELECT * FROM characters
WHERE village = ?
ORDER BY party, user;

-- name: ListCharactersFromParty :many
SELECT * FROM characters
WHERE party = ?
ORDER BY user;

-- name: ListCharactersFromUser :many
SELECT * FROM characters
WHERE user = ?
ORDER BY village;

-- name: CreateCharacter :one
INSERT INTO characters (
    user,
    party,
    village,
    gender,
    name,
    surname,
    age,
    portrait,
    background,
    physique,
    skin,
    hair,
    face,
    speech,
    clothing,
    virtue,
    vice,
    reputation,
    misfortune,
    description,
    hp,
    max_hp,
    str,
    max_str,
    dex,
    max_dex,
    will,
    max_will,
    inventory
) VALUES (
    ?, ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?,  ?
)
RETURNING *;

-- name: DeleteCharacter :exec
DELETE FROM characters
WHERE id = ?;

-- name: GetIdentity :one
SELECT characters.id, gender, name, surname, age, portrait, sqlc.embed(backgrounds) FROM characters
INNER JOIN backgrounds ON background = backgrounds.id
WHERE characters.id = ? LIMIT 1;

-- name: GetIdentityFromName :one
SELECT characters.id, gender, name, surname, age, portrait, sqlc.embed(backgrounds) FROM characters
INNER JOIN backgrounds ON background = backgrounds.id
WHERE characters.name = ? AND characters.surname = ? LIMIT 1;


-- name: UpdateIdentity :exec
UPDATE characters
set gender = ?,
name = ?,
surname = ?,
age = ?,
portrait = ?,
background = ?
WHERE id = ?;

-- name: ListAssociations :many
SELECT characters.id, sqlc.embed(users), sqlc.embed(parties), sqlc.embed(villages) FROM characters
INNER JOIN users ON user = users.id
INNER JOIN parties ON party = parties.id
INNER JOIN villages ON village = villages.id
ORDER BY village, party, user;

-- name: GetAssociations :one
SELECT characters.id, sqlc.embed(users), sqlc.embed(parties), sqlc.embed(villages) FROM characters
INNER JOIN users ON user = users.id
INNER JOIN parties ON party = parties.id
INNER JOIN villages ON village = villages.id
WHERE characters.id = ? LIMIT 1;

-- name: GetAssociationsFromName :one
SELECT characters.id, sqlc.embed(users), sqlc.embed(parties), sqlc.embed(villages) FROM characters
INNER JOIN users ON user = users.id
INNER JOIN parties ON party = parties.id
INNER JOIN villages ON village = villages.id
WHERE characters.name = ? AND characters.surname = ? LIMIT 1;

-- name: UpdateAssociations :exec
UPDATE characters
set user = ?,
party = ?,
village = ?
WHERE id = ?;

-- name: GetTraits :one
SELECT id, physique, skin, hair, face, speech, clothing, virtue, vice, reputation, misfortune FROM characters
WHERE id = ? LIMIT 1;

-- name: GetTraitsFromName :one
SELECT id, physique, skin, hair, face, speech, clothing, virtue, vice, reputation, misfortune FROM characters
WHERE name = ? AND surname = ? LIMIT 1;

-- name: UpdateTraits :exec
UPDATE characters
set physique = ?,
skin = ?,
hair = ?,
face = ?,
speech = ?,
clothing = ?,
virtue = ?,
vice = ?,
reputation = ?,
misfortune = ?
WHERE id = ?;

-- name: GetStatus :one
SELECT id, hp, max_hp, str, max_str, dex, max_dex, will, max_will FROM characters
WHERE id = ? LIMIT 1;

-- name: GetStatusFromName :one
SELECT id, hp, max_hp, str, max_str, dex, max_dex, will, max_will FROM characters
WHERE name = ? AND surname = ? LIMIT 1;

-- name: UpdateStatus :exec
UPDATE characters
set hp = ?,
str = ?,
dex = ?,
will = ?
WHERE id = ?;

-- name: UpdateMaxStatus :exec
UPDATE characters
set max_hp = ?,
max_str = ?,
max_dex = ?,
max_will = ?
WHERE id = ?;
