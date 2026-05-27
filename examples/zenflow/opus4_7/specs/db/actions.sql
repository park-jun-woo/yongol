CREATE TABLE actions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL DEFAULT 0 REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    config TEXT NOT NULL DEFAULT '',
    sequence_order BIGINT NOT NULL
);

CREATE INDEX idx_actions_workflow_id ON actions(workflow_id);
