-- name: OwnerLookupWebhook :one
SELECT org_id FROM webhooks WHERE id = @id;

-- name: WebhookFindByID :one
SELECT * FROM webhooks WHERE id = @id;

-- name: WebhookCreate :one
INSERT INTO webhooks (org_id, url, event_type)
VALUES (@org_id, @url, @event_type)
RETURNING *;

-- name: WebhookListByOrg :many
-- +no-pagination
SELECT * FROM webhooks WHERE org_id = @org_id ORDER BY created_at DESC;

-- name: WebhookDelete :exec
DELETE FROM webhooks WHERE id = @id;

-- name: WebhookListByOrgAndEvent :many
-- +no-pagination
SELECT * FROM webhooks WHERE org_id = @org_id AND event_type = @event_type;
