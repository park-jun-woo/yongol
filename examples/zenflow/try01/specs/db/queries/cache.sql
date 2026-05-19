-- name: CacheSet :exec
INSERT INTO fullend_cache (key, value, expires_at)
VALUES (@key, @value, @expires_at)
ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value, expires_at = EXCLUDED.expires_at;

-- name: CacheGet :one
SELECT value FROM fullend_cache WHERE key = @key AND expires_at > NOW();

-- name: CacheDelete :exec
DELETE FROM fullend_cache WHERE key = @key;
