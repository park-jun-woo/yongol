-- @archived
CREATE TABLE actions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    action_type TEXT NOT NULL,
    payload_template JSONB NOT NULL DEFAULT '{}',
    sequence_order BIGINT NOT NULL
);
