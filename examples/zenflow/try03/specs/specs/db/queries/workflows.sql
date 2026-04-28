-- name: WorkflowCreate :one
INSERT INTO workflows (org_id, title, trigger_event, status)
VALUES (@org_id, @title, @trigger_event, @status)
RETURNING *;

-- name: WorkflowFindByID :one
SELECT * FROM workflows
WHERE id = @id;

-- name: WorkflowListByOrgID :many
-- +no-pagination
SELECT * FROM workflows
WHERE org_id = @org_id
ORDER BY created_at DESC;

-- name: WorkflowUpdateStatus :exec
UPDATE workflows
SET status = @status
WHERE id = @id;

-- name: OwnerLookupWorkflow :one
SELECT org_id FROM workflows
WHERE id = @id;
