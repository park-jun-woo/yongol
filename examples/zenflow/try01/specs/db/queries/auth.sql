-- name: LoginLookup :one
SELECT id, password_hash, role, org_id, email
FROM users
WHERE email = @email;

-- name: RefreshTokenInsert :exec
INSERT INTO refresh_tokens (token_hash, claims, expires_at)
VALUES (@token_hash, @claims, @expires_at);

-- name: RefreshTokenFindByHash :one
SELECT claims, expires_at, revoked_at
FROM refresh_tokens
WHERE token_hash = @token_hash;

-- name: RefreshTokenRevoke :exec
UPDATE refresh_tokens
SET revoked_at = CURRENT_TIMESTAMP
WHERE token_hash = @token_hash AND revoked_at IS NULL;

-- name: RefreshTokenRevokeAll :exec
UPDATE refresh_tokens
SET revoked_at = CURRENT_TIMESTAMP
WHERE revoked_at IS NULL AND claims @> @matcher;
