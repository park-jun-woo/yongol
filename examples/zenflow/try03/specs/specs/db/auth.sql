-- @archived
CREATE TABLE refresh_tokens (
    token_hash  TEXT        PRIMARY KEY, -- @nosensitive
    claims      JSONB       NOT NULL,
    expires_at  TIMESTAMPTZ NOT NULL,
    revoked_at  TIMESTAMPTZ, -- @nullable
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX refresh_tokens_claims_idx ON refresh_tokens USING GIN (claims);
