CREATE TABLE execution_logs (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL DEFAULT 0 REFERENCES workflows(id),
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    status VARCHAR(32) NOT NULL DEFAULT 'success',
    credits_spent BIGINT NOT NULL DEFAULT 0,
    executed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX execution_logs_workflow_id_idx ON execution_logs(workflow_id);
CREATE INDEX execution_logs_org_id_idx ON execution_logs(org_id);
