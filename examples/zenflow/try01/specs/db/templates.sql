CREATE TABLE templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_workflow_id UUID NOT NULL REFERENCES workflows(id),
    org_id UUID NOT NULL REFERENCES organizations(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    category TEXT NOT NULL,
    clone_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_templates_source ON templates(source_workflow_id);
