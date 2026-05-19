CREATE TABLE actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000' REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    config TEXT NOT NULL DEFAULT '',
    sequence_order BIGINT NOT NULL
);

CREATE INDEX idx_actions_workflow_id ON actions(workflow_id);
