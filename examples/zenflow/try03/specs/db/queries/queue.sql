-- name: QueuePublish :exec
INSERT INTO fullend_queue (topic, payload, priority, deliver_at, traceparent)
VALUES (@topic, @payload, @priority, @deliver_at, @traceparent);

-- name: QueuePoll :one
UPDATE fullend_queue SET status = 'processing', updated_at = NOW()
WHERE id = (SELECT fq.id FROM fullend_queue fq WHERE fq.topic = @topic AND fq.status = 'pending' AND fq.deliver_at <= NOW() ORDER BY CASE fq.priority WHEN 'high' THEN 1 WHEN 'normal' THEN 2 WHEN 'low' THEN 3 ELSE 4 END ASC, fq.created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED)
RETURNING *;

-- name: QueueAck :exec
UPDATE fullend_queue SET status = 'done', updated_at = NOW() WHERE id = @id;
