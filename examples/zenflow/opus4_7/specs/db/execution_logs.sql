CREATE TABLE execution_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES workflows(id),
    org_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES organizations(id),
    status VARCHAR(20) NOT NULL DEFAULT '',
    credits_spent BIGINT NOT NULL DEFAULT 0,
    report_key VARCHAR(255) NOT NULL DEFAULT '',
    executed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_execution_logs_workflow_id ON execution_logs(workflow_id);
CREATE INDEX idx_execution_logs_org_id ON execution_logs(org_id);
