-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING *;

-- name: UserFindByEmail :one
-- +allow-sensitive password_hash is required by auth login path
SELECT * FROM users WHERE email = @email;
