CREATE TABLE execution_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id),
    org_id UUID NOT NULL REFERENCES organizations(id),
    status TEXT NOT NULL,
    credits_spent BIGINT NOT NULL DEFAULT 0,
    report_key TEXT, -- @nullable
    executed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_execution_logs_workflow_id ON execution_logs(workflow_id);
CREATE INDEX idx_execution_logs_org_id ON execution_logs(org_id);
