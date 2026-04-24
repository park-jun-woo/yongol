CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL, -- @sensitive
    role VARCHAR(32) NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- @sentinel
INSERT INTO users (id, email, password_hash, role)
OVERRIDING SYSTEM VALUE
VALUES (0, 'nobody@system', '', 'member')
ON CONFLICT DO NOTHING;
