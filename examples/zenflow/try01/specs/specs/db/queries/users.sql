-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING *;

-- name: UserFindByID :one
-- +allow-sensitive
SELECT * FROM users WHERE id = @id;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;
