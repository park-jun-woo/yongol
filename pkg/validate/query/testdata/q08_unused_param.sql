-- name: FindUserUnused :one
SELECT id, email FROM users WHERE id = @id;
