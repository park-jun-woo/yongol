-- name: UserCreate :one
-- +allow-sensitive
INSERT INTO users (email, password_hash)
VALUES (@email, @password_hash)
RETURNING *;

-- name: UserFindByID :one
-- +allow-sensitive
SELECT * FROM users WHERE id = @id;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;
