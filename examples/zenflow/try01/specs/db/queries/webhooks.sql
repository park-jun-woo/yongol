-- name: WebhookFindByID :one
SELECT id, org_id, url, event_type, created_at
FROM webhooks
WHERE id = @id;

-- name: WebhookListByOrgID :many
SELECT id, org_id, url, event_type, created_at
FROM webhooks
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT 1000;

-- name: WebhookCreate :one
INSERT INTO webhooks (org_id, url, event_type)
VALUES (@org_id, @url, @event_type)
RETURNING id, org_id, url, event_type, created_at;

-- name: WebhookDelete :exec
DELETE FROM webhooks
WHERE id = @id;
