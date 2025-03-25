-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserFromName :one
SELECT * FROM users
WHERE name = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
set name = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: ListParties :many
SELECT * FROM parties
ORDER BY name;

-- name: GetParty :one
SELECT * FROM parties
WHERE id = ? LIMIT 1;

-- name: GetPartyFromName :one
SELECT * FROM parties
WHERE name = ? LIMIT 1;

-- name: CreateParty :one
INSERT INTO parties (
    name, description, inventory
) VALUES (
    ?, ?, NULL
)
RETURNING *;

-- name: UpdateParty :exec
UPDATE parties
set name = ?
WHERE id = ?;

-- name: DeleteParty :exec
DELETE FROM parties
WHERE id = ?;

-- name: ListVillages :many
SELECT * FROM villages
ORDER BY name;

-- name: GetVillage :one
SELECT * FROM villages
WHERE id = ?;

-- name: GetVillageFromName :one
SELECT * FROM villages
WHERE name = ?;

-- name: CreateVillage :one
INSERT INTO villages (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: UpdateVillage :exec
UPDATE villages
set name = ?
WHERE id = ?;

-- name: DeleteVillage :exec
DELETE FROM villages
WHERE id = ?;
