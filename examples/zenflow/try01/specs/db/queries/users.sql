-- name: UserCreate :one
INSERT INTO users (email, password_hash, role)
VALUES (@email, @password_hash, @role)
RETURNING id, email, password_hash, role, created_at;

-- name: UserFindByEmail :one
SELECT id, email, password_hash, role, created_at FROM users WHERE email = @email;

-- name: UserFindByID :one
SELECT id, email, password_hash, role, created_at FROM users WHERE id = @id;

-- name: LoginLookup :one
SELECT id, password_hash, jsonb_build_object(
  'user_id', id,
  'email',   email,
  'role',    role
) AS claims
FROM users WHERE email = @email;
