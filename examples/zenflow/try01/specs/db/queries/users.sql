-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role, name)
VALUES (@org_id, @email, @password_hash, @role, @name)
RETURNING *;

-- name: UserFindByID :one
-- +allow-sensitive password_hash is required by auth middleware
SELECT * FROM users WHERE id = @id;

-- name: UserFindByEmail :one
-- +allow-sensitive password_hash is required by auth login path
SELECT * FROM users WHERE email = @email;
