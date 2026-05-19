-- name: UserFindByID :one
-- +allow-sensitive
SELECT * FROM users WHERE id = @id;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserListByOrgID :many
-- +no-pagination
SELECT id, org_id, email, role FROM users WHERE org_id = @org_id;

-- name: UserCountByOrgID :one
SELECT COUNT(*) FROM users WHERE org_id = @org_id;
