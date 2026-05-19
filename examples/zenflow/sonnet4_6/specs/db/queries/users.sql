-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING *;

-- name: UserFindByID :one
SELECT id, org_id, email, role FROM users WHERE id = @id;

-- name: UserListByOrg :many
-- +no-pagination
SELECT id, org_id, email, role FROM users WHERE org_id = @org_id;
