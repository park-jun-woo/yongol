-- name: UserCreate :one
INSERT INTO users (org_id, email, password_hash, role, claims)
VALUES (@org_id, @email, @password_hash, @role, @claims)
RETURNING id, org_id, email, role;

-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users
WHERE email = @email;

-- name: UserFindByID :one
SELECT id, org_id, email, role FROM users
WHERE id = @id;
