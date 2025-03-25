-- name: GetGender :one
SELECT * FROM genders
WHERE id = ? LIMIT 1;

-- name: GetRandomGender :one
SELECT * FROM genders
ORDER BY RANDOM() LIMIT 1;

-- name: ListGenders :many
SELECT * FROM genders
ORDER BY genders;

-- name: CreateGender :one
INSERT INTO genders (
    gender
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteGender :exec
DELETE FROM genders
WHERE id = ?;

-- name: GetBackground :one
SELECT * FROM backgrounds
WHERE id = ? LIMIT 1;

-- name: GetBackgroundFromTitle :one
SELECT * FROM backgrounds
WHERE title = ? LIMIT 1;

-- name: GetRandomBackground :one
SELECT * FROM backgrounds
ORDER BY RANDOM() LIMIT 1;

-- name: ListBackgrounds :many
SELECT * FROM backgrounds
ORDER BY title;

-- name: CreateBackground :one
INSERT INTO backgrounds (
    title, description, image
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: DeleteBackground :exec
DELETE FROM backgrounds
WHERE id = ?;

-- name: GetName :one
SELECT * FROM names
WHERE id = ? LIMIT 1;

-- name: GetRandomName :one
SELECT * FROM names
ORDER BY RANDOM() LIMIT 1;

-- name: ListNames :many
SELECT * FROM names
ORDER BY genders, name;

-- name: CreateName :one
INSERT INTO names (
    name
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteName :exec
DELETE FROM names
WHERE id = ?;

-- name: GetRandomMasculineName :one
SELECT names.* FROM name_genders
INNER JOIN genders ON name_genders.gender == genders.id
INNER JOIN names ON name_genders.name == names.id
WHERE genders.gender = "Male"
ORDER BY RANDOM() LIMIT 1;

-- name: GetRandomFeminineName :one
SELECT names.* FROM name_genders
INNER JOIN genders ON name_genders.genders == genders.id
INNER JOIN names ON name_genders.name == names.id
WHERE genders.gender = "Female"
ORDER BY RANDOM() LIMIT 1;

-- name: ListMasculineNames :many
SELECT names.name FROM name_genders
INNER JOIN genders ON name_genders.genders == genders.id
INNER JOIN names ON name_genders.name == names.id
WHERE genders.gender = "Male"
ORDER BY names.name;

-- name: ListFeminineNames :many
SELECT names.name FROM name_genders
INNER JOIN genders ON name_genders.genders == genders.id
INNER JOIN names ON name_genders.name == names.id
WHERE genders.gender = "Female"
ORDER BY names.name;

-- name: MakeNameMasculine :one
WITH g AS (SELECT * FROM genders WHERE gender = "Male"),
    n AS (SELECT * FROM names WHERE name = "?")
INSERT INTO name_genders (
    name, gender
) VALUES (
    n.id, g.id
)
RETURNING *;

-- name: MakeNameFeminine :one
WITH g AS (SELECT * FROM genders WHERE gender = "Female"),
    n AS (SELECT * FROM names WHERE name = "?")
INSERT INTO name_genders (
    name, gender
) VALUES (
    n.id, g.id
)
RETURNING *;
--SELECT name.name_id, genders.genders_id FROM name
--CROSS JOIN genders
--WHERE genders.genders = "Female" AND name.name = ?
--RETURNING *;

-- name: DeleteNameGender :exec
DELETE FROM name_genders
WHERE id = ?;

-- name: DeleteNameGenderByValue :exec
DELETE FROM name_genders
WHERE name = ? AND gender = ?;

-- name: GetSurname :one
SELECT * FROM surnames
WHERE id = ? LIMIT 1;

-- name: GetRandomSurname :one
SELECT * FROM surnames
ORDER BY RANDOM() LIMIT 1;

-- name: ListSurnames :many
SELECT * FROM surnames
ORDER BY surnames;

-- name: CreateSurname :one
INSERT INTO surnames (
    surname
) VALUES (
    ?
)
RETURNING *;

-- name: DeleteSurname :exec
DELETE FROM surnames
WHERE id = ?;
