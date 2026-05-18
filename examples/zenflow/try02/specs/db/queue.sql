-- @archived
CREATE TABLE fullend_queue (
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    topic        TEXT NOT NULL,
    payload      JSONB NOT NULL,
    priority     TEXT NOT NULL DEFAULT 'normal',
    status       TEXT NOT NULL DEFAULT 'pending',
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deliver_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMPTZ, -- @nullable
    traceparent  TEXT NOT NULL DEFAULT ''
);
CREATE INDEX idx_fullend_queue_pending
    ON fullend_queue (topic, status, deliver_at) WHERE status = 'pending';
