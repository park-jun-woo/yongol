-- @archived
CREATE TABLE fullend_sessions (
    key        TEXT PRIMARY KEY,
    value      BYTEA NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);
CREATE INDEX idx_fullend_sessions_expires ON fullend_sessions(expires_at);
