CREATE TABLE workflows (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'paused', 'archived')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- @sentinel
INSERT INTO workflows (id, org_id, title, trigger_event, status)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'system', 'system', 'draft')
ON CONFLICT DO NOTHING;

CREATE INDEX workflows_org_id_idx ON workflows (org_id);
