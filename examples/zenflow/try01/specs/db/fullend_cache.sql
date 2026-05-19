-- @archived
CREATE TABLE fullend_cache (
    key TEXT PRIMARY KEY,
    value BYTEA NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
