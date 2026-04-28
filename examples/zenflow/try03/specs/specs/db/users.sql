CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL, -- @sensitive
    role TEXT NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'member')),
    claims JSONB NOT NULL DEFAULT '{}' -- @sensitive
);
