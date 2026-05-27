CREATE TABLE workflows (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    version BIGINT NOT NULL DEFAULT 1,
    root_workflow_id BIGINT NOT NULL DEFAULT 0,
    assigned_to BIGINT NOT NULL DEFAULT 0,
    assignment_confidence VARCHAR(10) NOT NULL DEFAULT 'none',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_workflows_org_id ON workflows(org_id);
CREATE INDEX idx_workflows_status ON workflows(status);

-- @sentinel
INSERT INTO workflows (id, org_id, title, trigger_event, status)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, '', '', 'draft')
ON CONFLICT DO NOTHING;
