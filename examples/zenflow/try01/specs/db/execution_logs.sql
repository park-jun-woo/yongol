CREATE TABLE execution_logs (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL REFERENCES workflows(id),
    org_id BIGINT NOT NULL REFERENCES organizations(id),
    status VARCHAR(32) NOT NULL DEFAULT 'pending',
    credits_spent BIGINT NOT NULL DEFAULT 0,
    executed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_execution_logs_workflow ON execution_logs(workflow_id);
CREATE INDEX idx_execution_logs_org ON execution_logs(org_id);
