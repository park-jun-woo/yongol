CREATE TABLE webhooks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id BIGINT NOT NULL DEFAULT 0 REFERENCES organizations(id),
    url TEXT NOT NULL,
    event_type VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_webhooks_org_id ON webhooks(org_id);
