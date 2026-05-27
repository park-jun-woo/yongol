-- name: LoginLookup :one
SELECT id, password_hash, claims FROM users WHERE email = @email;

-- name: RefreshTokenInsert :exec
INSERT INTO refresh_tokens (token_hash, claims, expires_at)
VALUES (@token_hash, @claims, @expires_at);

-- name: RefreshTokenConsume :one
WITH consumed AS (
    UPDATE refresh_tokens SET revoked_at = NOW()
    WHERE token_hash = @token_hash AND revoked_at IS NULL AND expires_at > NOW()
    RETURNING claims
)
SELECT claims FROM consumed;

-- name: RefreshTokenCheckReuse :one
SELECT claims FROM refresh_tokens WHERE token_hash = @token_hash AND revoked_at IS NOT NULL;

-- name: RefreshTokenRevoke :exec
UPDATE refresh_tokens SET revoked_at = NOW()
WHERE token_hash = @token_hash AND revoked_at IS NULL;

-- name: RefreshTokenRevokeAll :exec
UPDATE refresh_tokens SET revoked_at = NOW()
WHERE revoked_at IS NULL AND claims @> @matcher;
