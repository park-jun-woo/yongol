CREATE TABLE execution_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id),
    org_id UUID NOT NULL REFERENCES organizations(id),
    status TEXT NOT NULL,
    credits_spent BIGINT NOT NULL DEFAULT 0,
    executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    report_key VARCHAR(255) NOT NULL DEFAULT ''
);
