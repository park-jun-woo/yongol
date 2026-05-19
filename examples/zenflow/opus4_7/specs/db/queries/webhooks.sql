-- name: WebhookCreate :one
INSERT INTO webhooks (org_id, url, event_type)
VALUES (@org_id, @url, @event_type)
RETURNING *;

-- name: WebhookListByOrgID :many
-- +no-pagination
SELECT * FROM webhooks WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WebhookFindByID :one
SELECT * FROM webhooks WHERE id = @id;

-- name: WebhookDelete :exec
DELETE FROM webhooks WHERE id = @id;

-- name: OwnerLookupWebhook :one
SELECT org_id FROM webhooks WHERE id = @id;
