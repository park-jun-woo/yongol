-- name: RefreshTokenInsert :exec
-- Upsert so a refresh token re-issued within the same second (the JWT carries
-- only second-precision exp and no jti, so rapid rotation can mint an
-- identical token_hash) updates the existing row and clears revoked_at instead
-- of failing on the primary-key conflict.
INSERT INTO refresh_tokens (token_hash, claims, expires_at)
VALUES (@token_hash, @claims, @expires_at)
ON CONFLICT (token_hash) DO UPDATE
SET claims = EXCLUDED.claims, expires_at = EXCLUDED.expires_at, revoked_at = NULL;

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
