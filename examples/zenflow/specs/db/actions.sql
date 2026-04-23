CREATE TABLE actions (
    id BIGSERIAL PRIMARY KEY,
    workflow_id BIGINT NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    payload_template JSONB NOT NULL DEFAULT '{}'::jsonb,
    sequence_order BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_actions_workflow_id ON actions(workflow_id);
