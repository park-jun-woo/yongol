-- name: UserFindByID :one
-- +allow-sensitive
SELECT * FROM users WHERE id = @id;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserListByOrg :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM users WHERE org_id = @org_id ORDER BY email ASC;

-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING *;
