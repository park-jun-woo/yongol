CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id),
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL, -- @sensitive
    role TEXT NOT NULL DEFAULT 'member'
);
