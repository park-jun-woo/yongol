-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role)
VALUES (@org_id, @email, @password_hash, @role)
RETURNING id, org_id, email, password_hash, role, created_at;

-- name: UserFindByEmail :one
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE email = @email;

-- name: UserFindByID :one
SELECT id, org_id, email, password_hash, role, created_at
FROM users
WHERE id = @id;
