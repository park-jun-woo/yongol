CREATE TABLE actions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    action_type VARCHAR(100) NOT NULL,
    payload_template TEXT NOT NULL DEFAULT '',
    sequence_order BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_actions_workflow ON actions(workflow_id);
