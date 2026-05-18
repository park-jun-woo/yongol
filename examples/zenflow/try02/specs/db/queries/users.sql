-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserFindByID :one
SELECT id, org_id, email, role FROM users WHERE id = @id;

-- name: UserCreate :one
INSERT INTO users (org_id, email, role, password_hash)
VALUES (@org_id, @email, @role, @password_hash)
RETURNING id, org_id, email, role;
