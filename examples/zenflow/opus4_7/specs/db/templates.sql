CREATE TABLE templates (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    source_workflow_id BIGINT NOT NULL DEFAULT 0 REFERENCES workflows(id),
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    category VARCHAR(100) NOT NULL DEFAULT '',
    clone_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_templates_source ON templates(source_workflow_id);
CREATE INDEX idx_templates_org_id ON templates(org_id);
