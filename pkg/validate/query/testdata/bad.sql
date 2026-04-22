-- name: DeleteAll :exec
DELETE FROM users;

-- name: UpdateAll :exec
UPDATE users SET active = false;

-- name: list_users :many
SELECT id FROM users LIMIT 10;

-- name: GetData :unknown
SELECT id FROM users;

-- name: ListAll :many
SELECT id FROM users;

-- name: ExecSelect :exec
SELECT id FROM users;
