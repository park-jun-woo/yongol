CREATE TABLE actions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    workflow_id BIGINT NOT NULL DEFAULT 0 REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    payload_template JSONB, -- @nullable
    sequence_order BIGINT NOT NULL
);
