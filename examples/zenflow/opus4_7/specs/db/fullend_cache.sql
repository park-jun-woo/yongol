-- @archived
CREATE TABLE fullend_cache (
    key TEXT PRIMARY KEY,
    value BYTEA NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX idx_fullend_cache_expires ON fullend_cache(expires_at);
