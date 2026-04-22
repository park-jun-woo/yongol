-- name: GetUserAll :one
SELECT * FROM users WHERE id = @id;
