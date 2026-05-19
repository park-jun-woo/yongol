CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES organizations(id),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL, -- @sensitive
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'member')),
    claims JSONB NOT NULL DEFAULT '{}', -- @sensitive
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- @sentinel
INSERT INTO users (id, org_id, email, password_hash, role, claims)
VALUES ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'nobody@system', '', 'member', '{}')
ON CONFLICT DO NOTHING;
