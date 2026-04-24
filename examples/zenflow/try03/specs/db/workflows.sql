CREATE TABLE workflows (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title VARCHAR(255) NOT NULL,
    trigger_event VARCHAR(255) NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'draft',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX workflows_org_id_idx ON workflows(org_id);

-- @sentinel
INSERT INTO workflows (id, org_id, title, trigger_event, status)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'system', 'none', 'draft')
ON CONFLICT DO NOTHING;
