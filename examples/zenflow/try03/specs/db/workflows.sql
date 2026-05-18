CREATE TABLE workflows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID NOT NULL REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'paused', 'archived')),
    version BIGINT NOT NULL DEFAULT 1,
    root_workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000',
    assigned_to TEXT NOT NULL DEFAULT '',
    assignment_confidence VARCHAR(10) NOT NULL DEFAULT 'none',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
