-- name: QueuePublish :exec
INSERT INTO fullend_queue (topic, payload, priority, deliver_at, traceparent)
VALUES (@topic, @payload, @priority, @deliver_at, @traceparent);

-- name: QueuePoll :many
SELECT id, topic, payload, traceparent
FROM fullend_queue
WHERE status = 'pending' AND deliver_at <= NOW()
ORDER BY
    CASE priority WHEN 'high' THEN 0 WHEN 'normal' THEN 1 ELSE 2 END,
    id
FOR UPDATE SKIP LOCKED
LIMIT 100;

-- name: QueueAck :exec
UPDATE fullend_queue SET status = @status, processed_at = NOW() WHERE id = @id;
