-- name: UserFindByEmail :one
-- +allow-sensitive
SELECT * FROM users WHERE email = @email;

-- name: UserFindByID :one
SELECT id, org_id, email, role FROM users WHERE id = @id;
