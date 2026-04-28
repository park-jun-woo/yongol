CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, -- @sensitive
    role TEXT NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member')),
    claims JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- @sentinel
INSERT INTO users (id, org_id, email, password_hash, role, claims)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'nobody@system', '', 'member', '{}'::jsonb)
ON CONFLICT DO NOTHING;
