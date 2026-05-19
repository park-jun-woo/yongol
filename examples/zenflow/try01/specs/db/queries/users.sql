-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserFindByID :one
-- +allow-sensitive
SELECT * FROM users WHERE id = @id;

-- name: UserListByOrg :many
-- +no-pagination
-- +allow-sensitive
SELECT * FROM users WHERE org_id = @org_id;
