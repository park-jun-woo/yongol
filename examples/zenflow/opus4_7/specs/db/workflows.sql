CREATE TABLE workflows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    version BIGINT NOT NULL DEFAULT 1,
    root_workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    assigned_to UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    assignment_confidence VARCHAR(10) NOT NULL DEFAULT 'none',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_workflows_org_id ON workflows(org_id);
CREATE INDEX idx_workflows_status ON workflows(status);
