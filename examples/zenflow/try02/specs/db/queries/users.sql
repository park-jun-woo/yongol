-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT id, org_id, email, password_hash, role, claims, created_at
FROM users WHERE email = @email;

-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role, claims)
VALUES (@org_id, @email, @password_hash, @role, @claims)
RETURNING id, org_id, email, role, claims, created_at;
