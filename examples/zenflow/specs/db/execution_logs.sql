CREATE TABLE execution_logs (
    id BIGSERIAL PRIMARY KEY,
    workflow_id BIGINT NOT NULL REFERENCES workflows(id),
    org_id BIGINT NOT NULL REFERENCES organizations(id),
    status TEXT NOT NULL,
    credits_spent BIGINT NOT NULL DEFAULT 0,
    executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_execution_logs_workflow_id ON execution_logs(workflow_id);
CREATE INDEX idx_execution_logs_org_id ON execution_logs(org_id);
