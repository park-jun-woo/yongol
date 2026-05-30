CREATE TABLE bans (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    gallery_id BIGINT NOT NULL REFERENCES galleries(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    reason TEXT NOT NULL DEFAULT '',
    expires_at TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- @sensitive
    UNIQUE (gallery_id, user_id)
);
