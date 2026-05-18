CREATE TABLE audit_logs (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    org_id UUID NOT NULL REFERENCES organizations(id),
    actor_id UUID NOT NULL REFERENCES users(id),
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(50) NOT NULL,
    resource_id BIGINT NOT NULL DEFAULT 0,
    detail TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX audit_logs_org_id_idx ON audit_logs (org_id);
CREATE INDEX audit_logs_created_at_idx ON audit_logs (created_at);
CREATE INDEX audit_logs_action_idx ON audit_logs (action);
