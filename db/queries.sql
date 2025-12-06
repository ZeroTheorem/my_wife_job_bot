-- name: CreateRow :exec
INSERT INTO notes (name, val, month, year)
VALUES (?, ?, ?, ?);

-- name: GetAvg :one
SELECT AVG(val) FROM notes
WHERE name = ? AND month = ? AND year = ?;

-- name: DeleteLastRow :one
DELETE FROM notes
WHERE id = (SELECT MAX(id) from notes)
RETURNING name, val;
