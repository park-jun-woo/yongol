-- @archived
CREATE TABLE refresh_tokens (
    token_hash TEXT PRIMARY KEY, -- @sensitive
    claims JSONB NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE, -- @nullable
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX refresh_tokens_claims_idx ON refresh_tokens USING GIN (claims);
