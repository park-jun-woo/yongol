-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event)
VALUES (@org_id, @title, @trigger_event)
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows
WHERE id = @id;

-- name: WorkflowListByOrgID :many
SELECT * FROM workflows
WHERE org_id = @org_id
ORDER BY created_at DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page)::int - 1) * sqlc.arg(per_page);

-- name: WorkflowCountByOrgID :one
SELECT COUNT(*) FROM workflows
WHERE org_id = @org_id;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows
SET status = @status
WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id AS owner_id FROM workflows
WHERE id = @resource_id;
