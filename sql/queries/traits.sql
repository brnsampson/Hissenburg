-- name: GetPhysique :one
SELECT * FROM physique
WHERE physique_id = ? LIMIT 1;

-- name: GetRandomPhysique :one
SELECT * FROM physique
ORDER BY RANDOM() LIMIT 1;

-- name: ListPhysiques :many
SELECT * FROM physique
ORDER BY physique;

-- name: CreatePhysique :one
INSERT INTO physique (
    physique
) VALUES (
    ?
)
RETURNING *;

-- name: DeletePhysique :exec
DELETE FROM physique
WHERE physique_id = ?;

-- name: GetSkin :one
SELECT * FROM skin
WHERE skin_id = ? LIMIT 1;

-- name: GetRandomSkin :one
SELECT * FROM skin
ORDER BY RANDOM() LIMIT 1;

-- name: ListSkin :many
SELECT * FROM skin
ORDER BY skin;

-- name: CreateSkin :one
INSERT INTO skin (
    skin
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteSkin :exec
DELETE FROM skin
WHERE skin_id = ?;

-- name: GetHair :one
SELECT * FROM hair
WHERE hair_id = ? LIMIT 1;

-- name: GetRandomHair :one
SELECT * FROM hair
ORDER BY RANDOM() LIMIT 1;

-- name: ListHair :many
SELECT * FROM hair
ORDER BY hair;

-- name: CreateHair :one
INSERT INTO hair (
    hair
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteHair :exec
DELETE FROM hair
WHERE hair_id = ?;

-- name: GetFace :one
SELECT * FROM face
WHERE face_id = ? LIMIT 1;

-- name: GetRandomFace :one
SELECT * FROM face
ORDER BY RANDOM() LIMIT 1;

-- name: ListFaces :many
SELECT * FROM face
ORDER BY face;

-- name: CreateFace :one
INSERT INTO face (
    face
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteFace :exec
DELETE FROM face
WHERE face_id = ?;

-- name: GetSpeech :one
SELECT * FROM speech
WHERE speech_id = ? LIMIT 1;

-- name: GetRandomSpeech :one
SELECT * FROM speech
ORDER BY RANDOM() LIMIT 1;

-- name: ListSpeech :many
SELECT * FROM speech
ORDER BY speech;

-- name: CreateSpeech :one
INSERT INTO speech (
    speech
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteSpeech :exec
DELETE FROM speech
WHERE speech_id = ?;

-- name: GetClothing :one
SELECT * FROM clothing
WHERE clothing_id = ? LIMIT 1;

-- name: GetRandomClothing :one
SELECT * FROM clothing
ORDER BY RANDOM() LIMIT 1;

-- name: ListClothing :many
SELECT * FROM clothing
ORDER BY clothing;

-- name: CreateClothing :one
INSERT INTO clothing (
    clothing
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteClothing :exec
DELETE FROM clothing
WHERE clothing_id = ?;

-- name: GetVirtue :one
SELECT * FROM virtue
WHERE virtue_id = ? LIMIT 1;

-- name: GetRandomVirtue :one
SELECT * FROM virtue
ORDER BY RANDOM() LIMIT 1;

-- name: ListVirtues :many
SELECT * FROM virtue
ORDER BY virtue;

-- name: CreateVirtue :one
INSERT INTO virtue (
    virtue
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteVirtue :exec
DELETE FROM virtue
WHERE virtue_id = ?;

-- name: GetVice :one
SELECT * FROM vice
WHERE vice_id = ? LIMIT 1;

-- name: GetRandomVice :one
SELECT * FROM vice
ORDER BY RANDOM() LIMIT 1;

-- name: ListVices :many
SELECT * FROM vice
ORDER BY vice;

-- name: CreateVice :one
INSERT INTO vice (
    vice
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteVice :exec
DELETE FROM vice
WHERE vice_id = ?;

-- name: GetReputation :one
SELECT * FROM reputation
WHERE reputation_id = ? LIMIT 1;

-- name: GetRandomReputation :one
SELECT * FROM reputation
ORDER BY RANDOM() LIMIT 1;

-- name: ListReputations :many
SELECT * FROM reputation
ORDER BY reputation;

-- name: CreateReputation :one
INSERT INTO reputation (
    reputation
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteReputation :exec
DELETE FROM reputation
WHERE reputation_id = ?;

-- name: GetMisfortune :one
SELECT * FROM misfortune
WHERE misfortune_id = ? LIMIT 1;

-- name: GetRandomMisfortune :one
SELECT * FROM misfortune
ORDER BY RANDOM() LIMIT 1;

-- name: ListMisfortunes :many
SELECT * FROM misfortune
ORDER BY misfortune;

-- name: CreateMisfortune :one
INSERT INTO misfortune (
    misfortune
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteMisfortune :exec
DELETE FROM misfortune
WHERE misfortune_id = ?;
