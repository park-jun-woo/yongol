CREATE TABLE actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    config TEXT NOT NULL DEFAULT '',
    sequence_order BIGINT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_actions_workflow_id ON actions(workflow_id);
CREATE INDEX idx_actions_sequence ON actions(workflow_id, sequence_order);
