CREATE TABLE workflows (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    title TEXT NOT NULL,
    trigger_event TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- @sentinel
INSERT INTO workflows (id, org_id, title, trigger_event, status)
OVERRIDING SYSTEM VALUE
VALUES (0, 0, 'system', 'system', 'draft')
ON CONFLICT DO NOTHING;
