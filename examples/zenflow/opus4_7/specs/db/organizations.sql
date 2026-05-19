CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    plan_type VARCHAR(20) NOT NULL DEFAULT 'free' CHECK (plan_type IN ('free', 'pro', 'enterprise')),
    credits_balance BIGINT NOT NULL DEFAULT 0,
    latitude TEXT NOT NULL DEFAULT '',
    longitude TEXT NOT NULL DEFAULT '',
    address_verified BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- @sentinel
INSERT INTO organizations (id, name, plan_type, credits_balance)
VALUES ('00000000-0000-0000-0000-000000000000', 'system', 'free', 0)
ON CONFLICT DO NOTHING;
