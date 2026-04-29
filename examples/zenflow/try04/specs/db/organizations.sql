CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    plan_type TEXT NOT NULL DEFAULT 'free',
    credits_balance BIGINT NOT NULL DEFAULT 0
);
