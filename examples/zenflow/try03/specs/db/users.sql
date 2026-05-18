CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL, -- @sensitive
    role VARCHAR(20) NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member'))
);
