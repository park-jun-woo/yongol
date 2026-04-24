CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL, -- @sensitive
    role VARCHAR(32) NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- @sentinel
INSERT INTO users (id, org_id, email, password_hash, role)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'nobody@system', '', 'member')
ON CONFLICT DO NOTHING;
