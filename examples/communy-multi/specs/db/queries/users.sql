-- name: UserCreate :one
INSERT INTO users (email, password_hash, nickname, role)
VALUES (@email, @password_hash, @nickname, @role)
RETURNING *;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserFindByID :one
SELECT id, email, nickname, role, created_at FROM users WHERE id = @id;
