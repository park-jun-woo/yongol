CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL DEFAULT '', -- @sensitive
    role VARCHAR(32) NOT NULL DEFAULT 'member',
    name VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- @sentinel
INSERT INTO users (id, org_id, email, password_hash, role, name)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'nobody@system', '', 'system', 'Nobody')
ON CONFLICT DO NOTHING;
