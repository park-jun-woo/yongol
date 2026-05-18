CREATE TABLE workflows (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'paused', 'archived')),
    version BIGINT NOT NULL DEFAULT 1,
    root_workflow_id BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- @sentinel
INSERT INTO workflows (id, org_id, title, trigger_event, status, version, root_workflow_id)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'system', 'none', 'draft', 1, 0)
ON CONFLICT DO NOTHING;
