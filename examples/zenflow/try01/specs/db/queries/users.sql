-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING id, org_id, email, password_hash, role, created_at;

-- name: UserFindByID :one
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE id = @id;

-- name: UserFindByEmail :one
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE email = @email;

-- name: UserListByOrgID :many
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT sqlc.arg(per_page) OFFSET sqlc.arg(page);

-- name: UserCountByOrgID :one
SELECT COUNT(*) as total
FROM users
WHERE org_id = @org_id;
