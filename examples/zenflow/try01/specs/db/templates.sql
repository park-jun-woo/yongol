CREATE TABLE templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id),
    title TEXT NOT NULL,
    description TEXT, -- @nullable
    category TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_templates_workflow_id ON templates(workflow_id);
CREATE INDEX idx_templates_category ON templates(category);
