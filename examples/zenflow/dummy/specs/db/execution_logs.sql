CREATE TABLE execution_logs (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL DEFAULT 0 REFERENCES workflows(id),
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    status TEXT NOT NULL DEFAULT 'success',
    credits_spent BIGINT NOT NULL DEFAULT 0,
    executed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
