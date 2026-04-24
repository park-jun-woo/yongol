-- name: RefreshTokenInsert :exec
INSERT INTO refresh_tokens (token_hash, claims, expires_at)
VALUES (@token_hash, @claims, @expires_at);

-- name: RefreshTokenFindByHash :one
SELECT claims, expires_at, revoked_at FROM refresh_tokens WHERE token_hash = @token_hash;

-- name: RefreshTokenRevoke :exec
UPDATE refresh_tokens SET revoked_at = NOW()
WHERE token_hash = @token_hash AND revoked_at IS NULL;

-- name: RefreshTokenRevokeAll :exec
UPDATE refresh_tokens SET revoked_at = NOW()
WHERE revoked_at IS NULL AND claims @> @matcher;
