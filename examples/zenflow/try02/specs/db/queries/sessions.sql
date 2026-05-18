-- name: SessionSet :exec
INSERT INTO fullend_sessions (key, value, expires_at) VALUES (@key, @value, @expires_at)
ON CONFLICT (key) DO UPDATE SET value = @value, expires_at = @expires_at;

-- name: SessionGet :one
SELECT value FROM fullend_sessions WHERE key = @key AND expires_at > NOW();

-- name: SessionDelete :exec
DELETE FROM fullend_sessions WHERE key = @key;
