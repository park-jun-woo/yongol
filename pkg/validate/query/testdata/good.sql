-- name: GetUser :one
SELECT id, email FROM users WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = @id;

-- name: UpdateUser :exec
UPDATE users SET email = @email WHERE id = @id;

-- name: ListUsers :many
SELECT id FROM users ORDER BY id DESC LIMIT @per_page;
