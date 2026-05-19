CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL, -- @sensitive
    role TEXT NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member'))
);
