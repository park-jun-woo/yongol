CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    plan_type VARCHAR(20) NOT NULL DEFAULT 'free' CHECK (plan_type IN ('free', 'pro', 'enterprise')),
    credits_balance BIGINT NOT NULL DEFAULT 0,
    address TEXT NOT NULL DEFAULT '',
    latitude TEXT NOT NULL DEFAULT '0',
    longitude TEXT NOT NULL DEFAULT '0',
    address_verified BOOLEAN NOT NULL DEFAULT false
);

-- @sentinel
INSERT INTO organizations (id, name, plan_type, credits_balance)
VALUES ('00000000-0000-0000-0000-000000000001', 'Seed Org', 'enterprise', 9999)
ON CONFLICT DO NOTHING;
