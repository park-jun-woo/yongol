CREATE TABLE execution_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    org_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT '',
    credits_spent BIGINT, -- @nullable
    report_file_key TEXT NOT NULL DEFAULT '',
    executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- @nullable
);
